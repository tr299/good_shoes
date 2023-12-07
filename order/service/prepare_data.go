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

func prepareDataToResponseGetSalesOrder(order *model_order.SalesOrderModel, orderItems []*model_order.SalesOrderItemModel) *model_order.SalesOrderDetail {
	data := &model_order.SalesOrderDetail{
		SalesOrder: *model_order.ConvertSalesOrderModelToSalesOrderResponse(order),
		Items:      nil,
	}

	return data
}
