package service

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"
    "good_shoes/common/model/model_product"
    "good_shoes/common/util"
    "good_shoes/logger"
    "good_shoes/product/repository"
    "gorm.io/gorm"
    "net/http"

    "good_shoes/common/config"
)

type Handler struct {
    config   *config.Config
    database *gorm.DB
    tracer   trace.Tracer
}

func NewHandler(config *config.Config, db *gorm.DB, tracer trace.Tracer) (*Handler, error) {
    return &Handler{
        config:   config,
        database: db,
        tracer:   tracer,
    }, nil
}

func (h *Handler) CreateProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "CreateProduct")
    defer span.End()

    req := &model_product.CreateProductRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter CreateProduct, data request = ", req)

    if err := validateCreateProduct(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    data, err := repo.CreateProduct(prepareDataToCreateProduct(req))
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
    }

    c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "UpdateProduct")
    defer span.End()

    req := &model_product.UpdateProductRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter UpdateProduct, data request = ", req)

    if err := validateUpdateProduct(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    data, err := repo.UpdateProduct(prepareDataToUpdateProduct(req))
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
    }

    c.JSON(http.StatusOK, data)
}

func (h *Handler) ListProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "ListProduct")
    defer span.End()

    req := &model_product.ListProductRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter GetProduct, data request = ", req)

    repo := repository.NewRepository(h.database)
    result, err := repo.ListProduct()
    if err != nil {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    data := prepareDataToResponseListProduct(result)
    c.JSON(http.StatusOK, data)
}

func (h *Handler) GetProduct(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "GetProduct")
    defer span.End()

    req := &model_product.GetProductByIdRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter GetProduct, data request = ", req)

    if err := validateGetProduct(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    data, err := repo.GetProductById(req.Id)
    if nil != err {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    response := model_product.ConvertProductModelToProductResponse(data)
    c.JSON(http.StatusOK, response)
}
