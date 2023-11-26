package service

import (
    "fmt"
    "github.com/gofrs/uuid"
    "good_shoes/common/model/model_product"
    "good_shoes/logger"
    "time"
)

func prepareDataToResponseListProduct(o []*model_product.ProductModel) *model_product.ListProductResponse {
    var items []model_product.ProductItem

    for _, p := range o {
        items = append(items, *model_product.ConvertProductModelToProductResponse(p))
    }

    return &model_product.ListProductResponse{Items: items}
}

func prepareDataToCreateProduct(req *model_product.CreateProductRequest) *model_product.ProductModel {
    // generate uuid
    uuid, err := uuid.NewV4()
    if err != nil {
        logger.Error(err)
        return nil
    }

    data := &model_product.ProductModel{
        Id:               fmt.Sprintf("PROD-%v", uuid),
        IsVariant:        req.IsVariant,
        Sku:              req.Sku,
        Name:             req.Name,
        Description:      req.Description,
        Description2:     req.Description2,
        Status:           req.Status,
        Type:             req.Type,
        OptionKey:        "",
        OptionValue:      "",
        Tags:             "",
        Price:            req.Price,
        SalePrice:        req.SalePrice,
        Cost:             req.Cost,
        CategoryIds:      req.Category,
        CheckInventory:   req.CheckInventory,
        MultipleVariants: req.MultipleVariant,
        TotalQuantity:    req.TotalQty,
        Brand:            req.Brand,
        ImageUrl:         req.ImageUrl,
    }

    createdAt := time.Now()
    data.CreatedAt = &createdAt

    return data
}

func prepareDataToUpdateProduct(req *model_product.UpdateProductRequest) *model_product.ProductModel {
    if nil == req {
        return nil
    }

    data := &model_product.ProductModel{
        Id:               req.Id,
        IsVariant:        req.IsVariant,
        Sku:              req.Sku,
        Name:             req.Name,
        Description:      req.Description,
        Description2:     req.Description2,
        Status:           req.Status,
        Type:             req.Type,
        OptionKey:        "",
        OptionValue:      "",
        Tags:             "",
        Price:            req.Price,
        SalePrice:        req.SalePrice,
        Cost:             req.Cost,
        CategoryIds:      req.Category,
        CheckInventory:   req.CheckInventory,
        MultipleVariants: req.MultipleVariant,
        TotalQuantity:    req.TotalQty,
        Brand:            req.Brand,
        ImageUrl:         req.ImageUrl,
    }

    updatedAt := time.Now()
    data.UpdatedAt = &updatedAt

    return data
}
