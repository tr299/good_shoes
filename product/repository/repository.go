package repository

import (
    "errors"
    "good_shoes/common/model/model_product"
    "good_shoes/logger"
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) CreateProduct(o *model_product.ProductModel) (*model_product.ProductModel, error) {
    if len(o.Id) == 0 {
        return nil, errors.New("product id is required")
    }

    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("products").Create(o).Error
    if err != nil {
        logger.Error("repository create product failed: ", err)
        return nil, err
    }

    return o, nil
}

func (r *Repository) UpdateProduct(o *model_product.ProductModel) (*model_product.ProductModel, error) {
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("products").Omit("id,created_at,deleted_at").Updates(o).Error
    if err != nil {
        logger.Error("repository update product failed: ", err)
        return nil, err
    }

    return o, nil
}

func (r *Repository) ListProduct() ([]*model_product.ProductModel, error) {
    var products []*model_product.ProductModel
    query := r.db.Session(&gorm.Session{NewDB: true})
    if err := query.Table("products").Find(&products).Error; err != nil {
        logger.Error("repository list product failed: ", err)
        return nil, err
    }

    return products, nil
}

func (r *Repository) GetProductById(id string) (*model_product.ProductModel, error) {
    var products *model_product.ProductModel
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("products").Where("id = ?", id).First(&products).Error
    if err != nil {
        logger.Error("repository list product failed: ", err)
        return nil, err
    }

    return products, nil
}
