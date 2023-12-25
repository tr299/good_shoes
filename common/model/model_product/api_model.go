package model_product

type ListProductRequest struct {
    InStockOnly bool    `form:"in_stock_only"`
    MinPrice    float64 `form:"min_price"`
    MaxPrice    float64 `form:"max_price"`
    Category    string  `form:"category"`
    Brand       string  `form:"brand"`
    Tag         string  `form:"tag"`
    ParentId    string  `form:"parent_id"`
    Page        int     `form:"page"`
    Limit       int     `form:"limit"`
}

type ListProductResponse struct {
    Items []ProductItem `json:"items"`
}

type OptionItem struct {
    Id       string `json:"id"`
    OptionId string `json:"option_id"`
    Label    string `json:"label"`
    Value    string `json:"value"`
    ImageUrl string `json:"image_url"`

    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt string `json:"deleted_at,omitempty"`
}

type Option struct {
    Id        string       `json:"id"`
    Name      string       `json:"name"`
    ProductId string       `json:"product_id"`
    Key       string       `json:"key"`
    Type      string       `json:"type"`
    Price     float64      `json:"price"`
    Items     []OptionItem `json:"items"`

    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt string `json:"deleted_at,omitempty"`
}

type ProductItem struct {
    Id              string   `json:"id"`
    ParentId        string   `json:"parent_id"`
    Sku             string   `json:"sku"`
    Name            string   `json:"name"`
    Status          string   `json:"status"`
    Type            string   `json:"type"`
    Price           float64  `json:"price"`
    Cost            float64  `json:"cost"`
    SalePrice       float64  `json:"sale_price"`
    TotalQty        uint32   `json:"total_quantity"`
    CheckInventory  bool     `json:"check_inventory"`
    MultipleVariant bool     `json:"multiple_variant"`
    Category        string   `json:"category"`
    Brand           string   `json:"brand"`
    CreatedAt       string   `json:"created_at"`
    UpdatedAt       string   `json:"updated_at"`
    DeletedAt       string   `json:"deleted_at"`
    Description     string   `json:"description"`
    Description2    string   `json:"description2"`
    ImageUrl        string   `json:"image_url"`
    IsVariant       bool     `json:"is_variant"`
    Tag             string   `json:"tag"`
    Options         []Option `json:"options"`
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
    Tag             string  `json:"tag"`
}

type GetProductByIdRequest struct {
    Id    string `uri:"id"`
    Size  string `form:"size"`
    Color string `form:"color"`
}

type DeleteProductByIdRequest struct {
    Id string `uri:"id"`
}

type CreateProductResponse struct {
    ProductId string `json:"product_id"`
    Message   string `json:"message"`
}
