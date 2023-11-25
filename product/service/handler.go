package service

import (
    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"

    "good_shoes/common/config"
)

type Handler struct {
    config *config.Config
    tracer trace.Tracer
}

func NewHandler(config *config.Config, tracer trace.Tracer) (*Handler, error) {
    return &Handler{
        config: config,
        tracer: tracer,
    }, nil
}

func (h *Handler) ListProduct(c *gin.Context) {

}

func (h *Handler) GetProduct(c *gin.Context) {

}
