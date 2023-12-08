package service

import (
    "errors"

    "good_shoes/common/model/model_order"
)

func validateUpdateOrderStatus(req *model_order.UpdateOrderStatusRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.Id) == 0 {
        return errors.New("order id is required")
    }

    if len(req.Status) == 0 {
        return errors.New("order status is required")
    }

    return nil
}
