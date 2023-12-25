package service

import (
    "errors"

    "good_shoes/common/model/model_inventory"
)

func validateAddInventory(req *model_inventory.AddInventoryRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.ProductId) == 0 {
        return errors.New("product id is required")
    }

    if req.Quantity == 0 {
        return errors.New("quantity is required")
    }

    return nil
}

func validateSubInventory(req *model_inventory.SubInventoryRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.ProductId) == 0 {
        return errors.New("product id is required")
    }

    if req.Quantity == 0 {
        return errors.New("quantity is required")
    }

    return nil
}
