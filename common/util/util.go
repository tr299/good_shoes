package util

import (
    "bytes"
    "context"
    "crypto/md5"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

    "good_shoes/logger"
)

func SendResponse(c *gin.Context, status int, data any) {
    c.JSON(status, data)
}

func DoHttpRequest(
    ctx context.Context,
    requestURI string,
    method string,
    queryParams map[string]string,
    body any) (int, []byte, error) {
    requestBody, err := json.Marshal(body)
    if err != nil {
        return http.StatusInternalServerError, nil, err
    }

    logger.Infof("====HTTP Request: %s %s", method, requestURI)
    logger.Infof("====Request: %s", requestBody)

    req, err := CreateHttpRequestWithContext(
        ctx,
        requestURI,
        method,
        queryParams,
        requestBody,
    )
    if err != nil {
        return http.StatusInternalServerError, nil, fmt.Errorf("create http request error: %w", err)
    }

    client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
    response, err := client.Do(req)
    if err != nil {
        return http.StatusInternalServerError, nil, err
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return http.StatusInternalServerError, nil, err
    }
    logger.Infof("====Response: %d %s", response.StatusCode, data)

    return response.StatusCode, data, err
}

func CreateHttpRequestWithContext(
    ctx context.Context,
    url string,
    method string,
    queryParams map[string]string,
    body []byte,
) (*http.Request, error) {
    var requestBuilder strings.Builder
    requestBuilder.WriteString(url)
    parameters := make(map[string]string)
    if queryParams != nil {
        c := "?"
        for key, val := range queryParams {
            requestBuilder.WriteString(fmt.Sprintf("%s%s=%s", c, key, val))
            c = "&"
            if val != "" {
                parameters[key] = val
            }
        }
    }
    bodyReader := bytes.NewReader(body)
    req, err := http.NewRequestWithContext(ctx, method, requestBuilder.String(), bodyReader)
    if err != nil {
        return nil, fmt.Errorf("can not create http request: %w", err)
    }

    signedHeaderMap := make(map[string]string)
    signedHeaderMap["Accept"] = "application/json"

    for k, v := range signedHeaderMap {
        req.Header.Add(k, v)
    }

    if len(body) > 0 {
        req.Header.Add("Content-Type", "application/json")
    }

    return req, nil
}

func NewIDFixLength(prefix string, length int) string {
    // Tính toán số lượng ký tự hexa cần sinh
    numHexChars := (length + 1) / 2
    randomBytes := make([]byte, numHexChars)
    rand.Read(randomBytes)

    // Thêm tgian để giảm tối đa trùng lặp
    currentTime := time.Now().UnixNano()
    dataToHash := append(randomBytes, []byte(fmt.Sprint(currentTime))...)
    hash := md5.Sum(dataToHash)
    md5String := hex.EncodeToString(hash[:])

    uniqueString := prefix + md5String
    uid := uuid.New()
    if len(uniqueString) > length {
        uniqueString = uniqueString[:length]
        id := base64.RawURLEncoding.EncodeToString(uid[:])[:len(uniqueString)/2]
        uniqueString = uniqueString[:len(uniqueString)/2] + id
    } else {
        id := base64.RawURLEncoding.EncodeToString(uid[:])
        uniqueString = uniqueString + id
        uniqueString = uniqueString[:length]
    }
    return uniqueString
}
