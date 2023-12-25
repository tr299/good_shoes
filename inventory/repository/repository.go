package repository

import (
    "gorm.io/gorm"

    "good_shoes/common/model/model_inventory"
    "good_shoes/logger"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) AddQty(req *model_inventory.AddInventoryRequest) error {
    query := r.db.Session(&gorm.Session{NewDB: true}).Table("products")
    query = query.Where("id = ?", req.ProductId)
    if len(req.ParentId) > 0 {
        query = query.Or("id = ?", req.ParentId)
    }
    err := query.Update("total_quantity", gorm.Expr("total_quantity + ?", req.Quantity)).Error
    if err != nil {
        logger.Error("repository list order failed: ", err)
        return err
    }

    return nil
}

func (r *Repository) SubQty(req *model_inventory.SubInventoryRequest) error {
    query := r.db.Session(&gorm.Session{NewDB: true}).Table("products")
    query = query.Where("id = ?", req.ProductId)
    if len(req.ParentId) > 0 {
        query = query.Or("id = ?", req.ParentId)
    }
    err := query.Update("total_quantity", gorm.Expr("total_quantity - ?", req.Quantity)).Error
    if err != nil {
        logger.Error("repository list order failed: ", err)
        return err
    }

    return nil
}
