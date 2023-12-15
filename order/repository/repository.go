package repository

import (
    "errors"
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
    offset := 0
    limit := 20

    if req.Page > 0 {
        offset = (req.Page - 1) * req.Limit
        limit = req.Limit
    }

    query := r.db.Session(&gorm.Session{NewDB: true}).Table("sales_orders")

    err := query.Limit(limit).Offset(offset).Find(&orders).Error
    if err != nil {
        logger.Error("repository list order failed: ", err)
        return nil, err
    }

    return orders, nil
}

func (r *Repository) GetOrderById(id string) (*model_order.SalesOrderModel, error) {
    var order *model_order.SalesOrderModel
    query := r.db.Session(&gorm.Session{NewDB: true}).Table("sales_orders")
    err := query.Table("sales_orders").Where("id = ?", id).First(&order).Error
    if err != nil {
        logger.Error("repository get order by id failed: ", err)
        return nil, err
    }

    return order, nil
}

func (r *Repository) ListOrderItems(orderId string) ([]*model_order.SalesOrderItemModel, error) {
    var items []*model_order.SalesOrderItemModel
    query := r.db.Session(&gorm.Session{NewDB: true}).Table("sales_order_items")

    err := query.Where("sales_order_id = ?", orderId).Find(&items).Error
    if err != nil {
        logger.Error("repository list order items failed: ", err)
        return nil, err
    }

    return items, nil
}

func (r *Repository) CreateSalesOrder(o *model_order.SalesOrderModel) (*model_order.SalesOrderModel, error) {
    if len(o.Id) == 0 {
        return nil, errors.New("product id is required")
    }

    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("sales_orders").Create(o).Error
    if err != nil {
        logger.Error("repository create sales order failed: ", err)
        return nil, err
    }

    return o, nil
}

func (r *Repository) CreateSalesOrderItems(o []model_order.SalesOrderItemModel) ([]model_order.SalesOrderItemModel, error) {
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("sales_order_items").Create(o).Error
    if err != nil {
        logger.Error("repository create sales order items failed: ", err)
        return nil, err
    }

    return o, nil
}

func (r *Repository) UpdateSalesOrderStatus(id, status string) (*model_order.SalesOrderModel, error) {
    query := r.db.Session(&gorm.Session{NewDB: true})
    err := query.Table("sales_orders").Where("id = ?", id).UpdateColumn("status", status).Error
    if err != nil {
        logger.Error("repository update order status failed: ", err)
        return nil, err
    }

    return r.GetOrderById(id)
}
