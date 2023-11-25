package cmd

import (
    "context"
    "net/http"
    "syscall"
    "time"

    "os"
    "os/signal"

    _ "github.com/lib/pq"
    "github.com/spf13/cobra"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/resource"
    tracesdk "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
    "good_shoes/common/config"
    "good_shoes/logger"
    "good_shoes/router"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

const (
    service = "good_shoes"
    version = "v1.0.0"
)

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Run Server",

    Run: runServerCommand,
}

func init() {
    rootCmd.AddCommand(serverCmd)
}

func runServerCommand(cmd *cobra.Command, args []string) {
    config, err := config.LoadConfig(".")
    if err != nil {
        logger.Fatal("cannot load config:", err)
    }

    logger.Init(logger.Config{
        Level:       config.LoggerConfig.Level,
        JSONFormat:  config.LoggerConfig.JSONFormat,
        EnableTrace: config.LoggerConfig.EnableTrace,
    })

    tp, err := initializeTracerProvider(config.Tracer)
    if err != nil {
        logger.Fatal(err)
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Cleanly shutdown and flush telemetry when the application exits.
    defer func(ctx context.Context) {
        // Do not make the application hang when it is shutdown.
        ctx, cancel = context.WithTimeout(ctx, time.Second*5)
        defer cancel()
        if err := tp.Shutdown(ctx); err != nil {
            logger.Fatal(err)
        }
    }(ctx)

    dbConnections := initializeDbConnection(config)
    server := initializeServer(config, dbConnections)

    logger.Debug("Start server")
    // Initializing the server in a goroutine so that
    // it won't block the graceful shutdown handling below
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatalf("listen: %s\n", err)
        }
    }()

    // Wait for interrupt signal to gracefully shutdown the server with
    // a timeout of 5 seconds.
    quit := make(chan os.Signal, 1)
    // kill (no param) default send syscall.SIGTERM
    // kill -2 is syscall.SIGINT
    // kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    logger.Info("Shutting down server...")

    // The context is used to inform the server it has 90 seconds to finish
    // the request it is currently handling
    ctx, cancel = context.WithTimeout(context.Background(), 90*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        logger.Fatal("Server forced to shutdown: ", err)
    }

    logger.Info("Server exiting")
}

func initializeServer(config config.Config, database *gorm.DB) *router.Server {
    server, err := router.NewServer(config, database)
    if err != nil {
        logger.Fatal("cannot create server:", err)
    }

    return server
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func initializeTracerProvider(config config.TracerConfig) (*tracesdk.TracerProvider, error) {
    // Create the Jaeger exporter
    exp, err := jaeger.New(
        jaeger.WithCollectorEndpoint(
            jaeger.WithEndpoint(config.Endpoint),
            jaeger.WithUsername(config.Username),
            jaeger.WithPassword(config.Password),
        ),
    )
    if err != nil {
        return nil, err
    }
    tp := tracesdk.NewTracerProvider(
        // Always be sure to batch in production.
        tracesdk.WithBatcher(exp),
        // Record information about this application in a Resource.
        tracesdk.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(service),
            semconv.ServiceVersionKey.String(version),
            attribute.String("environment", config.Environment),
        )),
    )

    // Register our TracerProvider as the global so any imported
    // instrumentation in the future will default to using it.
    otel.SetTracerProvider(tp)
    otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

    return tp, nil
}

func initializeDbConnection(c config.Config) *gorm.DB {
    logger.Info("start init database connections")

    DbConn, err := gorm.Open(sqlite.Open(c.Database.Source), &gorm.Config{})
    if err != nil {
        logger.Error("connect to token database failed => ", err)
    }

    logger.Info("connect to token database successfully!")

    return DbConn
}
