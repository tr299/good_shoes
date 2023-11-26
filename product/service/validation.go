package service

import (
    "errors"
    "good_shoes/common/model/model_product"
)

func validateCreateProduct(req *model_product.CreateProductRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.Name) == 0 {
        return errors.New("name is required")
    }

    return nil
}

func validateUpdateProduct(req *model_product.UpdateProductRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.Id) == 0 {
        return errors.New("product id is required")
    }

    return nil
}

func validateGetProduct(req *model_product.GetProductByIdRequest) error {
    if nil == req {
        return errors.New("request is required")
    }

    if len(req.Id) == 0 {
        return errors.New("product id is required")
    }

    return nil
}
