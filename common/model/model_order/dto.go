package model_order

import (
    "fmt"
    "time"

    "github.com/gofrs/uuid"
    "gorm.io/gorm"

    "good_shoes/logger"
)

type SalesOrderItemModel struct {
    Id              string
    SalesOrderID    string
    ProductID       string
    ParentProductID string
    Type            string
    Sku             string
    Name            string
    MediaUrl        string
    IsVirtual       bool
    IsGiftProduct   bool
    RequireShipping bool

    // Quantity
    QtyOrdered     float32
    QtyInvoiced    float32
    QtyFulfilled   float32
    QtyShipped     float32
    QtyRefunded    float32
    QtyRestocked   float32
    QtyCanceled    float32
    QtyBackordered float32

    // Price
    Price               float64
    PriceInclTax        float64
    CustomPrice         float64
    CustomPriceInclTax  float64
    SpecialPrice        float64
    SpecialPriceInclTax float64

    // Discount
    DiscountAmount         float64
    ShippingDiscountAmount float64
    AppliedDiscounts       string

    // Tax
    Taxable           bool
    TaxAmount         float64
    DiscountTaxAmount float64
    AppliedTaxes      string

    // shipping
    ShippingAmount float64

    // row total
    RowTotal        float64
    RowTotalInclTax float64

    Note     string
    MetaData string

    CreatedAt *time.Time
    UpdatedAt *time.Time
    DeletedAt *gorm.DeletedAt
}

type SalesOrderModel struct {
    Id          string
    OrderNumber string

    // status
    Status            string
    PaymentStatus     string
    FulfillmentStatus string
    ShipmentStatus    string
    ProcessStatus     string

    // total qty order_item
    TotalItemQty   float64
    TotalItemCount int64

    // customer
    CustomerID          string
    CustomerExternalID  string
    CustomerEmail       string
    CustomerPhone       string
    CustomerFirstName   string
    CustomerMiddleName  string
    CustomerLastName    string
    CustomerTaxNumber   string
    CustomerCompanyName string

    // currency
    Currency             string
    BaseCurrency         string
    CurrencyExchangeRate float64

    // subtotal
    Subtotal        float64
    SubtotalInclTax float64

    // discount
    DiscountPercent  float64
    DiscountAmount   float64
    AppliedDiscounts string

    // shipping
    ShippingMethod      string
    ShippingDescription string
    ShippingAmount      float64

    // tax
    TaxAmount float64

    // grand total
    ExtraAmount float64
    GrandTotal  float64

    // meta data and additional
    Note     string
    Tags     string
    MetaData string

    CreatedAt *time.Time
    UpdatedAt *time.Time
    DeletedAt *gorm.DeletedAt
}

func ConvertSalesOrderModelToSalesOrderResponse(o *SalesOrderModel) *SalesOrder {
    data := &SalesOrder{
        Id:                   o.Id,
        OrderNumber:          o.OrderNumber,
        Status:               o.Status,
        PaymentStatus:        o.PaymentStatus,
        FulfillmentStatus:    o.FulfillmentStatus,
        ShipmentStatus:       o.ShipmentStatus,
        ProcessStatus:        o.ProcessStatus,
        TotalItemQty:         o.TotalItemQty,
        TotalItemCount:       o.TotalItemCount,
        CustomerID:           o.CustomerID,
        CustomerExternalID:   o.CustomerExternalID,
        CustomerEmail:        o.CustomerEmail,
        CustomerPhone:        o.CustomerPhone,
        CustomerFirstName:    o.CustomerFirstName,
        CustomerMiddleName:   o.CustomerMiddleName,
        CustomerLastName:     o.CustomerLastName,
        CustomerTaxNumber:    o.CustomerTaxNumber,
        CustomerCompanyName:  o.CustomerCompanyName,
        Currency:             o.Currency,
        BaseCurrency:         o.BaseCurrency,
        CurrencyExchangeRate: o.CurrencyExchangeRate,
        Subtotal:             o.Subtotal,
        SubtotalInclTax:      o.SubtotalInclTax,
        DiscountPercent:      o.DiscountPercent,
        DiscountAmount:       o.DiscountAmount,
        AppliedDiscounts:     o.AppliedDiscounts,
        ShippingMethod:       o.ShippingMethod,
        ShippingDescription:  o.ShippingDescription,
        ShippingAmount:       o.ShippingAmount,
        TaxAmount:            o.TaxAmount,
        ExtraAmount:          o.ExtraAmount,
        GrandTotal:           o.GrandTotal,
        Note:                 o.Note,
        Tags:                 o.Tags,
        MetaData:             o.MetaData,
    }

    if nil != o.CreatedAt {
        data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
    }

    if nil != o.UpdatedAt {
        data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
    }

    return data
}

func ConvertOrderItemsModelToOrderItems(o *SalesOrderItemModel) *SalesOrderItem {
    data := &SalesOrderItem{
        Id:                     o.Id,
        SalesOrderID:           o.SalesOrderID,
        ProductID:              o.ProductID,
        ParentProductID:        o.ParentProductID,
        Type:                   o.Type,
        Sku:                    o.Sku,
        Name:                   o.Name,
        MediaUrl:               o.MediaUrl,
        IsVirtual:              o.IsVirtual,
        IsGiftProduct:          o.IsGiftProduct,
        RequireShipping:        o.RequireShipping,
        QtyOrdered:             o.QtyOrdered,
        QtyInvoiced:            o.QtyInvoiced,
        QtyFulfilled:           o.QtyFulfilled,
        QtyShipped:             o.QtyShipped,
        QtyRefunded:            o.QtyRefunded,
        QtyRestocked:           o.QtyRestocked,
        QtyCanceled:            o.QtyCanceled,
        QtyBackordered:         o.QtyBackordered,
        Price:                  o.Price,
        PriceInclTax:           o.PriceInclTax,
        CustomPrice:            o.CustomPrice,
        CustomPriceInclTax:     o.CustomPriceInclTax,
        SpecialPrice:           o.SpecialPrice,
        SpecialPriceInclTax:    o.SpecialPriceInclTax,
        DiscountAmount:         o.DiscountAmount,
        ShippingDiscountAmount: o.ShippingDiscountAmount,
        AppliedDiscounts:       o.AppliedDiscounts,
        Taxable:                o.Taxable,
        TaxAmount:              o.TaxAmount,
        DiscountTaxAmount:      o.DiscountTaxAmount,
        AppliedTaxes:           o.AppliedTaxes,
        ShippingAmount:         o.ShippingAmount,
        RowTotal:               o.RowTotal,
        RowTotalInclTax:        o.RowTotalInclTax,
        Note:                   o.Note,
        MetaData:               o.MetaData,
    }

    if nil != o.CreatedAt {
        data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
    }

    if nil != o.UpdatedAt {
        data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
    }

    return data
}

func ConverOrderItemReqToOrderItemModel(o *SalesOrderItem) *SalesOrderItemModel {
    uuid, err := uuid.NewV4()
    if err != nil {
        logger.Error(err)
        return nil
    }
    createdAt := time.Now()
    return &SalesOrderItemModel{
        Id:                     fmt.Sprintf("ORI-%v", uuid),
        SalesOrderID:           o.SalesOrderID,
        ProductID:              o.ProductID,
        ParentProductID:        o.ParentProductID,
        Type:                   o.Type,
        Sku:                    o.Sku,
        Name:                   o.Name,
        MediaUrl:               o.MediaUrl,
        IsVirtual:              o.IsVirtual,
        IsGiftProduct:          o.IsGiftProduct,
        RequireShipping:        o.RequireShipping,
        QtyOrdered:             o.QtyOrdered,
        QtyInvoiced:            o.QtyInvoiced,
        QtyFulfilled:           o.QtyFulfilled,
        QtyShipped:             o.QtyShipped,
        QtyRefunded:            o.QtyRefunded,
        QtyRestocked:           o.QtyRestocked,
        QtyCanceled:            o.QtyCanceled,
        QtyBackordered:         o.QtyBackordered,
        Price:                  o.Price,
        PriceInclTax:           o.PriceInclTax,
        CustomPrice:            o.CustomPrice,
        CustomPriceInclTax:     o.CustomPriceInclTax,
        SpecialPrice:           o.SpecialPrice,
        SpecialPriceInclTax:    o.SpecialPriceInclTax,
        DiscountAmount:         o.DiscountAmount,
        ShippingDiscountAmount: o.ShippingDiscountAmount,
        AppliedDiscounts:       o.AppliedDiscounts,
        Taxable:                o.Taxable,
        TaxAmount:              o.TaxAmount,
        DiscountTaxAmount:      o.DiscountTaxAmount,
        AppliedTaxes:           o.AppliedTaxes,
        ShippingAmount:         o.ShippingAmount,
        RowTotal:               o.RowTotal,
        RowTotalInclTax:        o.RowTotalInclTax,
        Note:                   o.Note,
        MetaData:               o.MetaData,
        CreatedAt:              &createdAt,
    }
}
