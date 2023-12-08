package service

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"
    "good_shoes/common/config"
    "good_shoes/logger"
    "net/http"
    "os"
    "path/filepath"
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

func (h *Handler) UploadFile(c *gin.Context) {
    uploadPath := "/var/www/html/uploads"

    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Tạo thư mục uploads nếu nó chưa tồn tại
    if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
        return
    }

    // Tạo tên tệp duy nhất cho ảnh
    filename := filepath.Join(uploadPath, file.Filename)

    // Lưu ảnh
    if err := c.SaveUploadedFile(file, filename); err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File '%s' uploaded successfully", file.Filename)})
}
