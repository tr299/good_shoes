package service

import (
    "errors"

    "github.com/tr299/good_shoes/common/model/model_wsp"
)

func validaCreateTokenRequest(req *model_wsp.CreateRegistrationTokenRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.MerchantID) == 0 {
        return errors.New("merchant_id is required")
    }

    if len(req.MerchantTokenRef) == 0 {
        return errors.New("merch_token_ref is required")
    }

    return nil
}

func validaGetTokenRequest(req *model_wsp.GetRegisterTokenRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.MerchantId) == 0 {
        return errors.New("merchant_id is required")
    }

    if len(req.RegistrationId) == 0 {
        return errors.New("merch_token_ref is required")
    }

    return nil
}
