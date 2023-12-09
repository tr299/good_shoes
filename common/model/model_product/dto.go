package model_product

import (
    "time"

    "gorm.io/gorm"
)

type OptionItemModel struct {
    Id       string
    OptionId string
    Label    string
    Value    string
    ImageUrl string

    CreatedAt *time.Time
    UpdatedAt *time.Time
    DeletedAt *gorm.DeletedAt
}

type ProductOptionModel struct {
    Id        string
    Name      string
    ProductId string
    Key       string
    Type      string
    Price     float64

    CreatedAt *time.Time
    UpdatedAt *time.Time
    DeletedAt *gorm.DeletedAt
}

type ProductModel struct {
    Id               string
    ParentId         string
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

    CreatedAt *time.Time
    UpdatedAt *time.Time
    DeletedAt *gorm.DeletedAt
}

func ConvertProductModelToProductResponse(o *ProductModel) *ProductItem {
    data := &ProductItem{
        Id:              o.Id,
        ParentId:        o.ParentId,
        IsVariant:       o.IsVariant,
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
        Tag:             o.Tags,
    }

    if nil != o.CreatedAt {
        data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
    }

    if nil != o.UpdatedAt {
        data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
    }

    return data
}

func ConvertOptionModelToOption(o *ProductOptionModel) Option {
    data := Option{
        Id:        o.Id,
        Name:      o.Name,
        ProductId: o.ProductId,
        Key:       o.Key,
        Type:      o.Type,
        Price:     o.Price,
        Items:     nil,
    }

    if nil != o.CreatedAt {
        data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
    }

    if nil != o.UpdatedAt {
        data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
    }

    return data
}

func ConvertOptionsModelToOptions(o []*ProductOptionModel) ([]Option, []string) {
    var data []Option
    var optionIds []string

    for _, model := range o {
        data = append(data, ConvertOptionModelToOption(model))
        optionIds = append(optionIds, model.Id)
    }

    return data, optionIds
}

func ConvertOptionItemModelToOptionItem(o *OptionItemModel) OptionItem {
    data := OptionItem{
        Id:       o.Id,
        OptionId: o.OptionId,
        Label:    o.Label,
        Value:    o.Value,
        ImageUrl: o.ImageUrl,
    }

    if nil != o.CreatedAt {
        data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
    }

    if nil != o.UpdatedAt {
        data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
    }

    return data
}

func ConvertOptionItemsModelToOptionItems(o []*OptionItemModel) []OptionItem {
    var data []OptionItem

    for _, model := range o {
        data = append(data, ConvertOptionItemModelToOptionItem(model))
    }

    return data
}
