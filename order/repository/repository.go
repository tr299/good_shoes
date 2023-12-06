package repository

import (
    "good_shoes/common/model/model_order"
    "good_shoes/logger"
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) ListOrder(req *model_order.ListSalesOrderRequest) ([]*model_order.SalesOrderModel, error) {
    var orders []*model_order.SalesOrderModel
    query := r.db.Session(&gorm.Session{NewDB: true}).Table("sales_orders")

    err := query.Debug().Find(&orders).Error
    if err != nil {
        logger.Error("repository list order failed: ", err)
        return nil, err
    }

    return orders, nil
}
