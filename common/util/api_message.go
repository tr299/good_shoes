package util

import (
    "context"

    "good_shoes/logger"

    "github.com/gin-gonic/gin"
)

const (
    INVALID_AUTHORIZATION        = "apimessage.invalid_authorization"
    INVALID_AUTHORIZATION_FORMAT = "apimessage.invalid_authorization_format"
    BAD_REQUEST                  = "apimessage.bad_request"
    TOKEN_NOT_FOUND              = "apimessage.token_not_found"
    TOKEN_EXPIRED                = "apimessage.token_expired"
    ERROR                        = "apimessage.error"
    INCORRECT_TERMINAL_SERIAL    = "apimessage.incorrect_terminal_serial"
    INVOICE_NOT_FOUND            = "apimessage.invoice_not_found"
    INVOICE_HIS_NOT_FOUND        = "apimessage.invoice_his_not_found"
    INVALID_PASSCODE             = "apimessage.invalid_passcode"
    MOBILE_BLOCK_SIGNIN          = "apimessage.mobile_block"
    PAYMENT_NOT_FOUND            = "apimessage.payment_not_found"
)

const (
    LanguageHeaderKey = "Lang"
)

func GetMessageMap() map[string]ErrorMessage {
    return map[string]ErrorMessage{

        BAD_REQUEST: {
            Status:      400,
            Description: "Invalid parameter.",
        },

        ERROR: {
            Status:      500,
            Description: "Internal server error",
        },
        INVALID_AUTHORIZATION: {
            Status:      1000,
            Description: "Authorization header is not provided.",
        },
        INVALID_AUTHORIZATION_FORMAT: {
            Status:      1001,
            Description: "invalid authorization header format.",
        },
        INCORRECT_TERMINAL_SERIAL: {
            Status:      1002,
            Description: "Terminal not allowed to connect, please verify its credentials",
        },
        INVOICE_NOT_FOUND: {
            Status:      1003,
            Description: "Invoice not exists.",
        },
        INVOICE_HIS_NOT_FOUND: {
            Status:      1004,
            Description: "Invoice not exists.",
        },
        INVALID_PASSCODE: {
            Status:      1005,
            Description: "Invalid passcode.",
        },
        MOBILE_BLOCK_SIGNIN: {
            Status:      1006,
            Description: "%d times wrong login infomation",
        },
        PAYMENT_NOT_FOUND: {
            Status:      1007,
            Description: "Payment not exists.",
        },
        TOKEN_NOT_FOUND: {
            Status:      1008,
            Description: "A working session has expired. Please sign in again.",
        },
        TOKEN_EXPIRED: {
            Status:      1009,
            Description: "A working session has expired. Please sign in again.",
        },
    }
}

type ErrorMessage struct {
    Status      int    `json:"Status"`
    Description string `json:"StatusDescription"`
}

func BindRequest(ctx context.Context, c *gin.Context, req interface{}) error {
    if err := c.ShouldBindUri(req); err != nil {
        logger.ErrorfContext(ctx, "parse request uri error:%s", err.Error())
        return err
    }

    if err := c.ShouldBind(req); err != nil {
        logger.ErrorfContext(ctx, "parse request body error:%s", err.Error())
        return err
    }
    return nil
}
