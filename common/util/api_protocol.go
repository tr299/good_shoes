package util

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "net/http/httptest"
    "net/url"
    "regexp"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"

    "logger"
)

type Protocol interface {
    RequestApi(ctx context.Context, url, method string, queryParams map[string]string, requestBody any) (int, []byte, error)
}

type RequestProtocol struct {
    Protocol
}

type RequestProtocolModules struct {
    TokenModule        *RequestProtocol
    InvoiceModule      *RequestProtocol
    PspSacombankModule *RequestProtocol
}

type HttpProtocol struct{}

type CallFuncProtocol struct{}

var mapRouter map[string]map[string]gin.HandlerFunc

func NewRequestProtocol(protocol string) *RequestProtocol {
    r := &RequestProtocol{}

    switch protocol {
    case HTTP_PROTOCOL:
        r.Protocol = NewHttpProtocol()
    case FUNC_CALL:
        r.Protocol = NewCallFuncProtocol()
    default:
        r.Protocol = NewHttpProtocol()
    }

    return r
}

func NewHttpProtocol() *HttpProtocol {
    return &HttpProtocol{}
}

func NewCallFuncProtocol() *CallFuncProtocol {
    return &CallFuncProtocol{}
}

func (http *HttpProtocol) RequestApi(ctx context.Context, url, method string, queryParams map[string]string, requestBody any) (int, []byte, error) {
    return DoHttpRequest(ctx, url, method, queryParams, requestBody)
}

func (f *CallFuncProtocol) RequestApi(ctx context.Context, url, method string, queryParams map[string]string, requestBody any) (int, []byte, error) {
    r, err := json.Marshal(requestBody)
    if nil != err {
        logger.Error(logrus.Fields{
            "package": "common",
            "func":    "RequestApi",
            "data":    requestBody,
            "err":     err,
            "message": "marshal json failed",
        })
        return http.StatusInternalServerError, nil, err
    }

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    req, err := CreateHttpRequestWithContext(ctx, url, method, queryParams, r)
    if nil != err {
        logger.Error(logrus.Fields{
            "package": "common",
            "func":    "RequestApi",
            "data":    fmt.Sprintf("url: %v, method: %v, queryParams: %v, r: %v", url, method, queryParams, r),
            "err":     err,
            "message": "create http request with context failed",
        })
        return http.StatusInternalServerError, nil, err
    }

    funcHandler, ginParams, err := getMapRouter(url, method)
    if nil == funcHandler || nil != err {
        logg, _ := json.Marshal(logrus.Fields{
            "package":     "common",
            "func":        "RequestApi",
            "url":         url,
            "method":      method,
            "queryParams": queryParams,
            "requestBody": requestBody,
            "err":         err,
            "message":     "can not get func wsp-go from map router",
        })
        logger.Error(string(logg))
        return http.StatusInternalServerError, nil, err
    }

    c.Request = req
    c.Params = ginParams
    funcHandler(c)

    b := make([]byte, w.Body.Len())
    w.Body.Read(b)
    responseBody := b

    return w.Code, responseBody, nil
}

func SetMapRouter(m map[string]map[string]gin.HandlerFunc) {
    mapRouter = m
}

func getMapRouter(link, method string) (gin.HandlerFunc, gin.Params, error) {
    ginParams := gin.Params{}
    u, err := url.Parse(link)
    if err != nil {
        logger.Error(logrus.Fields{
            "package": "common",
            "func":    "getMapRouter",
            "link":    link,
            "err":     err,
            "message": "link does not parse",
        })
        return nil, ginParams, err
    }

    for s, m := range mapRouter {
        isMatching, mapGinParams := MatchingDynamicPath(s, u.Path)
        if !isMatching {
            continue
        }

        // convert map[string]string to gin.Params
        for key, value := range mapGinParams {
            ginParams = append(ginParams, gin.Param{Key: key, Value: value})
        }

        return m[method], ginParams, nil
    }

    logger.Error(logrus.Fields{
        "package":   "common",
        "func":      "getMapRouter",
        "user_path": u.Path,
        "message":   "path does not found",
    })

    return nil, ginParams, errors.New("path not found")
}

func MatchingDynamicPath(routerPath, userPath string) (bool, map[string]string) {
    replaceRouterPathPattern := `:[^/]+`
    userPathRegexPattern := `(.+)`
    regexpRouterPath := regexp.MustCompile(replaceRouterPathPattern)
    replacedRouterPath := regexpRouterPath.ReplaceAllString(routerPath, userPathRegexPattern)
    pathParamKeys := regexpRouterPath.FindAllString(routerPath, -1)
    regex := regexp.MustCompile(replacedRouterPath)
    matches := regex.FindStringSubmatch(userPath)
    mapGinParams := map[string]string{}

    if len(matches) > len(pathParamKeys) {
        for i, s := range pathParamKeys {
            key := strings.ReplaceAll(s, ":", "")
            mapGinParams[key] = matches[i+1]
        }
        return true, mapGinParams
    }

    //URL does not match
    return false, mapGinParams
}
