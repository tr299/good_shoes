package model_inventory

type AddInventoryRequest struct {
    ProductId string `json:"product_id"`
    ParentId  string `json:"parent_id"`
    Quantity  int    `json:"quantity"`
}

type SubInventoryRequest struct {
    ProductId string `json:"product_id"`
    ParentId  string `json:"parent_id"`
    Quantity  int    `json:"quantity"`
}
