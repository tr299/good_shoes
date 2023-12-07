package router

import (
    "gorm.io/gorm"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
    "golang.org/x/net/context"

    "good_shoes/common/config"
    orderService "good_shoes/order/service"
    productService "good_shoes/product/service"
)

type Server struct {
    config   config.Config
    database *gorm.DB
    srv      *http.Server
    router   *gin.Engine
    //receiptMaker ReceiptMaker
    lstLoginImie map[string]LoginImie
}

type LoginImie struct {
    IMIE     string
    IP       string
    Number   int
    LastDate time.Time
}

// Create one tracer per package
// NOTE: You only need a tracer if you are creating your own spans
var tracer trace.Tracer

func init() {
    // Name the tracer after the package, or the service if you are in main
    tracer = otel.Tracer("router")
}

// NewServer creates new server instance
func NewServer(config config.Config, database *gorm.DB) (*Server, error) {
    server := &Server{
        config:       config,
        database:     database,
        lstLoginImie: make(map[string]LoginImie),
    }

    //server.receiptMaker = receiptMaker
    server.setupRoute()

    server.srv = &http.Server{
        Addr:    config.ServerAddress,
        Handler: server.router,
    }

    return server, nil
}

func (server *Server) ListenAndServe() error {
    return server.srv.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
    return server.srv.Shutdown(ctx)
}

func (server *Server) setupRoute() {
    //gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.Use(otelgin.Middleware("router"))
    apiPrefix := server.config.ApiPrefix

    // add jwt middleware
    authMiddleware, err := initJwt()
    if err != nil {
        log.Fatal("JWT Error:" + err.Error())
    }
    if errInit := authMiddleware.MiddlewareInit(); nil != errInit {
        log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
    }

    router.POST("/login", authMiddleware.LoginHandler)

    authRouter := router.Group(apiPrefix)
    authRouter.Use(authMiddleware.MiddlewareFunc())

    // product module
    productHandler, _ := productService.NewHandler(&server.config, server.database, tracer)
    authRouter.POST("/v1/products", productHandler.CreateProduct)
    authRouter.PUT("/v1/products/:id", productHandler.UpdateProduct)
    authRouter.GET("/v1/products", productHandler.ListProduct)
    authRouter.GET("/v1/products/:id", productHandler.GetProduct)

    // sales order module
    orderHandler, _ := orderService.NewHandler(&server.config, server.database, tracer)
    authRouter.POST("/v1/orders", orderHandler.CreateSalesOrder)
    authRouter.PUT("/v1/orders/:id", orderHandler.UpdateSalesOrder)
    authRouter.GET("/v1/orders", orderHandler.ListSalesOrder)
    authRouter.GET("/v1/orders/:id", orderHandler.GetSalesOrder)

    server.router = router
}
