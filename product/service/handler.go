package service

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel/trace"

    "good_shoes/common/config"
    "good_shoes/common/model/model_wsp"
    "good_shoes/common/util"
    "good_shoes/logger"
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

func (h *Handler) CreateRegisterToken(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "CreateToken")
    defer span.End()

    req := &model_wsp.CreateRegistrationTokenRequest{}
    //langCode := c.MustGet(languageHeaderKey).(string)

    if err := util.BindRequest(ctx, c, req); nil != err {
        //result := db.GetErrorMessage(ctx, h.store, common.BAD_REQUEST, "langCode")
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter CreateToken, data request = ", req)

    if err := validaCreateTokenRequest(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    // request POST to msp-go
    mspCreateTokenUrl := fmt.Sprintf("%v/ddsp/api/v1/merchants/%v/dd_tokens/%v", h.config.MspGoConfig.BaseUrl, req.MerchantID, req.MerchantTokenRef)
    statusCode, mspResponse, err := util.DoHttpRequest(ctx, mspCreateTokenUrl, "POST", nil, req)
    if err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, err)
        return
    }

    response := &model_wsp.CreateRegistrationTokenResponse{}
    if err := json.Unmarshal(mspResponse, &response); err != nil {
        logger.ErrorContext(ctx, err.Error())
        c.JSON(http.StatusInternalServerError, response)
    }

    c.JSON(statusCode, response)
}

func (h *Handler) GetRegisterToken(c *gin.Context) {
    newCtx := context.WithValue(c.Request.Context(), "Lang", c.Request.Header.Get(util.LanguageHeaderKey))
    ctx, span := h.tracer.Start(newCtx, "GetToken")
    defer span.End()

    req := &model_wsp.GetRegisterTokenRequest{}
    //langCode := c.MustGet(languageHeaderKey).(string)

    if err := util.BindRequest(ctx, c, req); nil != err {
        //result := db.GetErrorMessage(ctx, h.store, common.BAD_REQUEST, "langCode")
        c.JSON(http.StatusBadRequest, err)
        return
    }

    logger.Infof("Enter GetToken, data request = ", req)

    if err := validaGetTokenRequest(req); nil != err {
        logger.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    // request GET to msp-go
    mspGetTokenUrl := fmt.Sprintf("%v/ddsp/api/v1/merchants/%v/dd_tokens/%v", h.config.MspGoConfig.BaseUrl, req.MerchantId, req.RegistrationId)
    statusCode, mspResponse, err := util.DoHttpRequest(ctx, mspGetTokenUrl, "GET", nil, nil)
    if err != nil {
        logger.Error(err)
        c.JSON(http.StatusInternalServerError, err)
        return
    }

    response := &model_wsp.GetRegisterTokenResponse{}
    if err := json.Unmarshal(mspResponse, &response); err != nil {
        logger.ErrorContext(ctx, err.Error())
        c.JSON(http.StatusInternalServerError, response)
    }

    c.JSON(statusCode, response)
}
