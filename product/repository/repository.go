package repository

import (
    "good_shoes/common/model/model_product"
    "good_shoes/logger"
    "gorm.io/gorm"
)

type Repo struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repo {
    return &Repo{db: db}
}

func (r *Repo) ListProduct() []model_product.ProductModel {
    var products []model_product.ProductModel
    query := r.db.Session(&gorm.Session{NewDB: true})
    if err := query.Table("products").Find(&products).Error; err != nil {
        logger.Error(err)
    }

    return products
}
