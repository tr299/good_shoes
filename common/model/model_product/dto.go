package model_product

import (
    "time"
)

type ProductModel struct {
    Id               string
    IsVariant        bool
    Sku              string
    Name             string
    Description      string
    Description2     string
    Status           string
    Barcode          string
    Type             string
    OptionKey        string
    OptionValue      string
    Position         string
    Tags             string
    Price            float64
    SalePrice        float64
    Cost             float64
    CategoryIds      string
    CheckInventory   bool
    MultipleVariants bool
    TotalQuantity    uint32
    Brand            string
    ImageUrl         string
    CreatedAt        *time.Time
    UpdatedAt        *time.Time
    DeletedAt        *time.Time
}

func ConvertProductModelToProductResponse(o *ProductModel) *ProductItem {
    data := &ProductItem{
        Id:              o.Id,
        Sku:             o.Sku,
        Name:            o.Name,
        Status:          o.Status,
        Type:            o.Type,
        Price:           o.Price,
        Cost:            o.Cost,
        SalePrice:       o.SalePrice,
        TotalQty:        o.TotalQuantity,
        CheckInventory:  o.CheckInventory,
        MultipleVariant: o.MultipleVariants,
        Category:        o.CategoryIds,
        Brand:           o.Brand,
        Description:     o.Description,
        Description2:    o.Description2,
        ImageUrl:        o.ImageUrl,
    }

    if nil != o.CreatedAt {
        data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
    }

    if nil != o.UpdatedAt {
        data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
    }

    if nil != o.DeletedAt {
        data.DeletedAt = o.DeletedAt.Format(time.RFC3339)
    }

    return data
}
