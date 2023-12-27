package service

import (
    "fmt"
    "strings"
    "time"

    "github.com/gofrs/uuid"

    "good_shoes/common/model/model_product"
    "good_shoes/logger"
)

func prepareDataToResponseListProduct(o []*model_product.ProductModel, totalProduct int64) *model_product.ListProductResponse {
    var items []model_product.ProductItem

    for _, p := range o {
        items = append(items, *model_product.ConvertProductModelToProductResponse(p))
    }

    return &model_product.ListProductResponse{
        Total: totalProduct,
        Items: items,
    }
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
        IsVariant:        false,
        Sku:              req.Sku,
        Name:             req.Name,
        Description:      req.Description,
        Description2:     req.Description2,
        Status:           req.Status,
        Type:             "simple",
        OptionKey:        "",
        OptionValue:      "",
        Tags:             req.Tag,
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

    if len(req.Options) > 0 {
        data.Type = "configurable"
    }

    createdAt := time.Now()
    data.CreatedAt = &createdAt

    return data
}

func prepareDataToCreateVariant(req *model_product.CreateProductRequest, variantName, parentId string, variantNorrameSortOrder []string) *model_product.ProductModel {
    // generate uuid
    uuid, err := uuid.NewV4()
    if err != nil {
        logger.Error(err)
        return nil
    }

    res := strings.Split(variantName, "-")

    data := &model_product.ProductModel{
        Id:               fmt.Sprintf("PROD-%v", uuid),
        ParentId:         parentId,
        IsVariant:        true,
        Sku:              fmt.Sprintf("%v-%v", req.Sku, variantName),
        Name:             fmt.Sprintf("%v-%v", req.Name, variantName),
        Description:      req.Description,
        Description2:     req.Description2,
        Status:           req.Status,
        Type:             "variant",
        OptionKey:        "", // Lưu giá trị size
        OptionValue:      "", // Lưu giá trị color
        Tags:             req.Tag,
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

    if len(res) == len(variantNorrameSortOrder) {
        for i, s := range variantNorrameSortOrder {
            if strings.EqualFold(s, "size") {
                data.OptionKey = res[i]
            }
            if strings.EqualFold(s, "color") {
                data.OptionValue = res[i]
            }
        }
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
        Tags:             req.Tag,
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

func prepareDataToCreateOptionItems(req []model_product.OptionItem, optionId string) []model_product.OptionItemModel {
    var data []model_product.OptionItemModel

    for _, item := range req {
        // generate uuid
        uuid, err := uuid.NewV4()
        if err != nil {
            logger.Error(err)
            return nil
        }
        createdAt := time.Now()
        data = append(data, model_product.OptionItemModel{
            Id:        fmt.Sprintf("ITEM-%v", uuid),
            OptionId:  optionId,
            Label:     item.Label,
            Value:     item.Value,
            ImageUrl:  item.ImageUrl,
            CreatedAt: &createdAt,
        })
    }

    return data
}

func prepareDataToCreateOptions(req []model_product.Option, productId string) []*model_product.ProductOptionModel {
    var data []*model_product.ProductOptionModel

    for _, option := range req {
        // generate uuid
        uuid, err := uuid.NewV4()
        if err != nil {
            logger.Error(err)
            return nil
        }
        createdAt := time.Now()
        optionId := fmt.Sprintf("OP-%v", uuid)
        o := &model_product.ProductOptionModel{
            Id:        optionId,
            Name:      option.Name,
            ProductId: productId,
            Key:       option.Key,
            Type:      option.Type,
            Price:     option.Price,
            CreatedAt: &createdAt,
        }
        data = append(data, o)
    }

    return data
}

func prepareOptionToResponse(options []model_product.Option, optionItems []model_product.OptionItem) []model_product.Option {
    var data []model_product.Option
    mapOptionIdToItems := map[string][]model_product.OptionItem{}
    for _, item := range optionItems {
        mapOptionIdToItems[item.OptionId] = append(mapOptionIdToItems[item.OptionId], item)
    }

    for _, option := range options {
        option.Items = append(option.Items, mapOptionIdToItems[option.Id]...)
        data = append(data, option)
    }

    return data
}

func prepareDataToCreateProductVariant(s1, s2 []string) []string {
    var result []string
    for _, n := range s1 {
        if len(s2) == 0 {
            return s1
        }

        for _, s := range s2 {
            result = append(result, fmt.Sprintf("%v-%v", s, n))
        }
    }

    return result
}
