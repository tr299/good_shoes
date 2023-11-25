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
    Price            string
    SalePrice        string
    Cost             string
    CategoryIds      string
    CheckInventory   bool
    MultipleVariants bool
    TotalQuantity    string
    Brand            string
    ImageUrl         string
    CreatedAt        *time.Time
    UpdatedAt        *time.Time
    DeletedAt        *time.Time
}

func ConvertProductModelToProductResponse(o *ProductModel) *ProductItem {
    return &ProductItem{
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
        CreatedAt:       o.CreatedAt.String(),
        UpdatedAt:       o.UpdatedAt.String(),
        DeletedAt:       o.DeletedAt.String(),
        Description:     o.Description,
        Description2:    o.Description2,
        ImageUrl:        o.ImageUrl,
    }
}
