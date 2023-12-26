package service

import (
    "context"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"
    "gorm.io/gorm"

    "good_shoes/common/config"
    "good_shoes/common/model/model_inventory"
    "good_shoes/common/util"
    "good_shoes/inventory/repository"
    "good_shoes/logger"
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

func (h *Handler) AddHandler(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "Add")
    defer span.End()

    req := &model_inventory.AddInventoryRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter Add inventory, data request = ", req)

    if err := validateAddInventory(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    err := repo.AddQty(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Update quantity success",
    })
}

func (h *Handler) SubHandler(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "Sub")
    defer span.End()

    req := &model_inventory.SubInventoryRequest{}

    if err := util.BindRequest(ctx, c, req); nil != err {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter Sub inventory, data request = ", req)

    if err := validateSubInventory(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    repo := repository.NewRepository(h.database)
    err := repo.SubQty(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Update quantity success",
    })
}
