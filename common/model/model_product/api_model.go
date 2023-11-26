package model_product

type ListProductRequest struct {
    InStockOnly bool    `form:"in_stock_only"`
    MinPrice    float64 `form:"min_price"`
    MaxPrice    float64 `form:"max_price"`
    Category    string  `form:"category"`
    Brand       string  `form:"brand"`
}

type ListProductResponse struct {
    Items []ProductItem `json:"items"`
}

type ProductItem struct {
    Id              string  `json:"id"`
    Sku             string  `json:"sku"`
    Name            string  `json:"name"`
    Status          string  `json:"status"`
    Type            string  `json:"type"`
    Price           float64 `json:"price"`
    Cost            float64 `json:"cost"`
    SalePrice       float64 `json:"sale_price"`
    TotalQty        uint32  `json:"total_quantity"`
    CheckInventory  bool    `json:"check_inventory"`
    MultipleVariant bool    `json:"multiple_variant"`
    Category        string  `json:"category"`
    Brand           string  `json:"brand"`
    CreatedAt       string  `json:"created_at"`
    UpdatedAt       string  `json:"updated_at"`
    DeletedAt       string  `json:"deleted_at"`
    Description     string  `json:"description"`
    Description2    string  `json:"description2"`
    ImageUrl        string  `json:"image_url"`
    IsVariant       bool    `json:"is_variant"`
}

type CreateProductRequest struct {
    ProductItem
}

type UpdateProductRequest struct {
    Id              string  `uri:"id"`
    Sku             string  `json:"sku"`
    Name            string  `json:"name"`
    Status          string  `json:"status"`
    Type            string  `json:"type"`
    Price           float64 `json:"price"`
    Cost            float64 `json:"cost"`
    SalePrice       float64 `json:"sale_price"`
    TotalQty        uint32  `json:"total_quantity"`
    CheckInventory  bool    `json:"check_inventory"`
    MultipleVariant bool    `json:"multiple_variant"`
    Category        string  `json:"category"`
    Brand           string  `json:"brand"`
    CreatedAt       string  `json:"created_at"`
    UpdatedAt       string  `json:"updated_at"`
    DeletedAt       string  `json:"deleted_at"`
    Description     string  `json:"description"`
    Description2    string  `json:"description2"`
    ImageUrl        string  `json:"image_url"`
    IsVariant       bool    `json:"is_variant"`
}

type GetProductByIdRequest struct {
    Id string `uri:"id"`
}
