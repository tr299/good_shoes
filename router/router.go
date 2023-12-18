package router

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
    "golang.org/x/net/context"
    "gorm.io/gorm"

    "good_shoes/common/config"
    mediaService "good_shoes/media/service"
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

    // setup CORS
    config := cors.Config{
        AllowAllOrigins: true,
        AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
    }
    router.Use(cors.New(config))

    // add jwt middleware
    authMiddleware, err := initJwt()
    if err != nil {
        log.Fatal("JWT Error:" + err.Error())
    }
    if errInit := authMiddleware.MiddlewareInit(); nil != errInit {
        log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
    }

    router.POST(apiPrefix+"/v1/login", authMiddleware.LoginHandler)

    authRouter := router.Group("")
    authRouter.Use(authMiddleware.MiddlewareFunc())

    // user
    authRouter.GET(apiPrefix+"/v1/user", func(c *gin.Context) {
        user := c.Query("user")

        if user == "admin" {
            c.JSON(http.StatusOK, gin.H{
                "name":  "Admin",
                "email": "admin@good-shoes.tr29.store",
                "image": "https://www.good-shoes.tr29.store/uploads/1702918213333424600_user.png",
            })
            return
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "message": "User not found",
        })
    })

    // product module
    productHandler, _ := productService.NewHandler(&server.config, server.database, tracer)
    router.GET(apiPrefix+"/v1/products", productHandler.ListProduct)
    router.GET(apiPrefix+"/v1/products/:id", productHandler.GetProduct)
    authRouter.POST(apiPrefix+"/v1/products", productHandler.CreateProduct)
    authRouter.PUT(apiPrefix+"/v1/products/:id", productHandler.UpdateProduct)
    authRouter.DELETE(apiPrefix+"/v1/products/:id", productHandler.DeleteProduct)

    // sales order module
    orderHandler, _ := orderService.NewHandler(&server.config, server.database, tracer)
    router.POST(apiPrefix+"/v1/orders", orderHandler.CreateSalesOrder)
    router.GET(apiPrefix+"/v1/orders", orderHandler.ListSalesOrder)
    router.GET(apiPrefix+"/v1/orders/:id", orderHandler.GetSalesOrder)
    authRouter.PUT(apiPrefix+"/v1/orders/:id/status", orderHandler.UpdateSalesOrderStatus)

    // upload image
    mediaHandler, _ := mediaService.NewHandler(&server.config, tracer)
    authRouter.POST(apiPrefix+"/v1/uploads", mediaHandler.UploadMultipleFile)

    server.router = router
}
