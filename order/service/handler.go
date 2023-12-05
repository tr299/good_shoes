package service

import (
    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"
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

func (h *Handler) CreateSalesOrder(c *gin.Context) {

}

func (h *Handler) UpdateSalesOrder(c *gin.Context) {

}

func (h *Handler) ListSalesOrder(c *gin.Context) {

}

func (h *Handler) GetSalesOrder(c *gin.Context) {

}
