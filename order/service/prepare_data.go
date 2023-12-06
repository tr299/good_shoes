package service

import (
    "good_shoes/common/model/model_order"
)

func prepareDataToResponseListSalesOrder(o []*model_order.SalesOrderModel) *model_order.ListSalesOrderResponse {
    var items []model_order.SalesOrder

    for _, p := range o {
        items = append(items, *model_order.ConvertSalesOrderModelToSalesOrderResponse(p))
    }

    return &model_order.ListSalesOrderResponse{Items: items}
}
