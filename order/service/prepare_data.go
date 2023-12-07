package service

import (
    "fmt"
    "github.com/gofrs/uuid"
    "good_shoes/common/model/model_order"
    "good_shoes/logger"
    "time"
)

func prepareDataToResponseListSalesOrder(o []*model_order.SalesOrderModel) *model_order.ListSalesOrderResponse {
    var items []model_order.SalesOrder

    for _, p := range o {
        items = append(items, *model_order.ConvertSalesOrderModelToSalesOrderResponse(p))
    }

    return &model_order.ListSalesOrderResponse{Items: items}
}

func prepareDataToResponseGetSalesOrder(order *model_order.SalesOrderModel, orderItems []*model_order.SalesOrderItemModel) *model_order.SalesOrderDetail {
    if nil == order {
        return nil
    }

    var items []model_order.SalesOrderItem

    for _, item := range orderItems {
        items = append(items, *model_order.ConvertOrderItemsModelToOrderItems(item))
    }

    data := &model_order.SalesOrderDetail{
        SalesOrder: *model_order.ConvertSalesOrderModelToSalesOrderResponse(order),
        Items:      items,
    }

    return data
}

func prepareDataToCreateSalesOrder(req *model_order.CreateSalesOrderRequest) *model_order.SalesOrderModel {
    // generate uuid
    uuid, err := uuid.NewV4()
    if err != nil {
        logger.Error(err)
        return nil
    }

    data := &model_order.SalesOrderModel{
        Id:                   fmt.Sprintf("ORD-%v", uuid),
        OrderNumber:          req.OrderNumber,
        Status:               req.Status,
        PaymentStatus:        req.PaymentStatus,
        FulfillmentStatus:    req.FulfillmentStatus,
        ShipmentStatus:       req.ShipmentStatus,
        ProcessStatus:        req.ProcessStatus,
        TotalItemQty:         req.TotalItemQty,
        TotalItemCount:       req.TotalItemCount,
        CustomerID:           req.CustomerID,
        CustomerExternalID:   req.CustomerExternalID,
        CustomerEmail:        req.CustomerEmail,
        CustomerPhone:        req.CustomerPhone,
        CustomerFirstName:    req.CustomerFirstName,
        CustomerMiddleName:   req.CustomerMiddleName,
        CustomerLastName:     req.CustomerLastName,
        CustomerTaxNumber:    req.CustomerTaxNumber,
        CustomerCompanyName:  req.CustomerCompanyName,
        Currency:             req.Currency,
        BaseCurrency:         req.BaseCurrency,
        CurrencyExchangeRate: req.CurrencyExchangeRate,
        Subtotal:             req.Subtotal,
        SubtotalInclTax:      req.SubtotalInclTax,
        DiscountPercent:      req.DiscountPercent,
        DiscountAmount:       req.DiscountAmount,
        AppliedDiscounts:     req.AppliedDiscounts,
        ShippingMethod:       req.ShippingMethod,
        ShippingDescription:  req.ShippingDescription,
        ShippingAmount:       req.ShippingAmount,
        TaxAmount:            req.TaxAmount,
        ExtraAmount:          req.ExtraAmount,
        GrandTotal:           req.GrandTotal,
        Note:                 req.Note,
        Tags:                 req.Tags,
        MetaData:             req.MetaData,
    }

    createdAt := time.Now()
    data.CreatedAt = &createdAt

    return data
}

func prepareDataToCreateSalesOrderItems(r []model_order.SalesOrderItem, orderId string) []model_order.SalesOrderItemModel {
    var data []model_order.SalesOrderItemModel

    for _, item := range r {
        item.SalesOrderID = orderId
        data = append(data, *model_order.ConverOrderItemReqToOrderItemModel(&item))
    }

    return data
}
