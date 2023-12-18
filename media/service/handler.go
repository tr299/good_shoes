package service

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "time"

    "good_shoes/common/config"
    "good_shoes/logger"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"
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
    newFileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
    filename := filepath.Join(uploadPath, newFileName)

    // Lưu ảnh
    if err := c.SaveUploadedFile(file, filename); err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "url":     "https://www.good-shoes.tr29.store/uploads/" + newFileName,
        "message": fmt.Sprintf("File '%s' uploaded successfully", file.Filename),
    })
}

func (h *Handler) UploadMultipleFile(c *gin.Context) {
    uploadPath := "/var/www/html/uploads"
    var urlResponses []string
    // Chấp nhận tất cả các tệp tin từ yêu cầu
    form, err := c.MultipartForm()
    if err != nil {
        c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
        return
    }

    // Lấy danh sách các tệp tin tải lên
    files := form.File["file"]

    // Tạo thư mục nếu nó chưa tồn tại
    err = os.MkdirAll(uploadPath, os.ModePerm)
    if err != nil {
        c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
        return
    }

    // Lặp qua từng tệp tin và lưu vào thư mục "uploads"
    for _, file := range files {
        // Mở tệp tin tải lên
        fileHandle, err := file.Open()
        defer fileHandle.Close()
        if err != nil {
            c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
            return
        }

        // Tạo tên tệp duy nhất cho ảnh
        fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename)
        regExp := regexp.MustCompile("[^a-zA-Z0-9.()]+")
        uniqueFileName := regExp.ReplaceAllString(fileName, "_")

        // Tạo tệp tin trên ổ đĩa và ghi dữ liệu vào nó
        dst, err := os.Create("/var/www/html/uploads/" + uniqueFileName)
        defer dst.Close()
        if err != nil {
            c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
            return
        }

        // Sao chép dữ liệu từ tệp tin tải lên vào tệp tin trên ổ đĩa
        _, err = io.Copy(dst, fileHandle)
        if err != nil {
            c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
            return
        }

        urlResponses = append(urlResponses, "https://www.good-shoes.tr29.store/uploads/"+uniqueFileName)
    }

    c.JSON(http.StatusOK, gin.H{
        "url":     urlResponses,
        "message": fmt.Sprintf("upload files successfully"),
    })
}
