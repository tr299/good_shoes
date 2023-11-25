package service

import (
    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"
    "good_shoes/logger"
    "good_shoes/product/repository"
    "gorm.io/gorm"

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

func (h *Handler) ListProduct(c *gin.Context) {
    repo := repository.NewRepository(h.database)
    result := repo.ListProduct()
    logger.Info("1")
    logger.Info(result)
    logger.Info("2")
}

func (h *Handler) GetProduct(c *gin.Context) {

}
