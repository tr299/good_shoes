package router

import (
    "gorm.io/gorm"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
    "golang.org/x/net/context"

    "good_shoes/common/config"
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
    tracer = otel.Tracer("git.onepay.vn/onepay/go/ddsp/router")
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

    // product router
    productConfig := server.config.ProductConfig
    productHandler, _ := productService.NewHandler(&server.config, server.database, tracer)
    router.POST(productConfig.ApiPrefix+"/v1/products", productHandler.CreateProduct)
    router.PUT(productConfig.ApiPrefix+"/v1/products/:id", productHandler.UpdateProduct)
    router.GET(productConfig.ApiPrefix+"/v1/products", productHandler.ListProduct)
    router.GET(productConfig.ApiPrefix+"/v1/products/:id", productHandler.GetProduct)

    server.router = router
}
