package model_order

import (
    "gorm.io/gorm"
    "time"
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
