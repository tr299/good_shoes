package repository

import (
    "errors"
    "time"

    "gorm.io/gorm"

    "good_shoes/common/model/model_product"
    "good_shoes/logger"
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

func (r *Repository) DeleteProduct(id string) error {
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("products").Where("id = ?", id).UpdateColumn("deleted_at", time.Now()).Error
    if err != nil {
        logger.Error("repository delete product failed: ", err)
        return err
    }

    return nil
}

func (r *Repository) ListProduct(req *model_product.ListProductRequest) ([]*model_product.ProductModel, error) {
    var products []*model_product.ProductModel
    query := r.db.Session(&gorm.Session{NewDB: true}).Table("products")
    offset := 0
    limit := 20
    // build filter
    if req.Page > 0 {
        offset = (req.Page - 1) * req.Limit
        limit = req.Limit
    }

    if req.InStockOnly {
        query = query.Where("total_quantity > 0")
    }

    if len(req.ParentId) > 0 {
        query = query.Where("parent_id = ?", req.ParentId)
    }

    if len(req.Brand) > 0 {
        query = query.Where("brand = ?", req.Brand)
    }

    if len(req.Tag) > 0 {
        query = query.Where("tags = ?", req.Tag)
    }

    if len(req.Category) > 0 {
        query = query.Where("category_ids = ?", req.Category)
    }

    if req.MaxPrice > 0 {
        query = query.Where("price <= ?", req.MaxPrice)
    }

    if req.MinPrice > 0 {
        query = query.Where("price >= ?", req.MinPrice)
    }

    err := query.Limit(limit).Offset(offset).Find(&products).Error
    if err != nil {
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

func (r *Repository) CreateProductOptions(o []*model_product.ProductOptionModel) ([]*model_product.ProductOptionModel, error) {
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("options").Create(o).Error
    if err != nil {
        logger.Error("repository create product options failed: ", err)
        return nil, err
    }

    return o, nil
}

func (r *Repository) CreateProductOptionItems(o []model_product.OptionItemModel) ([]model_product.OptionItemModel, error) {
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("option_items").Create(o).Error
    if err != nil {
        logger.Error("repository create option items failed: ", err)
        return nil, err
    }

    return o, nil
}

func (r *Repository) ListProductOptionsByProductId(productId string) ([]*model_product.ProductOptionModel, error) {
    var options []*model_product.ProductOptionModel
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("options").Where("product_id = ?", productId).Find(&options).Error
    if err != nil {
        logger.Error("repository list product options failed: ", err)
        return nil, err
    }

    return options, nil
}

func (r *Repository) ListProductOptionItemsByOptionIds(optionIds []string) ([]*model_product.OptionItemModel, error) {
    var optionItems []*model_product.OptionItemModel
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("option_items").Where("option_id in ?", optionIds).Find(&optionItems).Error
    if err != nil {
        logger.Error("repository list product option items failed: ", err)
        return nil, err
    }

    return optionItems, nil
}
