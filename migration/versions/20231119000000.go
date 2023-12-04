package versions

import (
    "gorm.io/gorm"
    "time"
)

func Version20231119000000(tx *gorm.DB) error {
    type SalesOrderItem struct {
        Id string `gorm:"TYPE:VARCHAR(100);NOT NULL;PRIMARY_KEY"`

        // general
        SalesOrderID    string `gorm:"TYPE:VARCHAR(100)"`
        ProductID       string `gorm:"TYPE:VARCHAR(100)"`
        ParentProductID string `gorm:"TYPE:VARCHAR(100)"`
        Type            string `gorm:"TYPE:VARCHAR(255)"`
        Sku             string `gorm:"TYPE:VARCHAR(255)"`
        Name            string `gorm:"TYPE:VARCHAR(255)"`
        MediaUrl        string `gorm:"TYPE:VARCHAR(255)"`
        IsVirtual       bool
        IsGiftProduct   bool
        RequireShipping bool

        // Quantity
        QtyOrdered     float32 `gorm:"TYPE:DECIMAL(12,4)"`
        QtyInvoiced    float32 `gorm:"TYPE:DECIMAL(12,4)"`
        QtyFulfilled   float32 `gorm:"TYPE:DECIMAL(12,4)"`
        QtyShipped     float32 `gorm:"TYPE:DECIMAL(12,4)"`
        QtyRefunded    float32 `gorm:"TYPE:DECIMAL(12,4)"`
        QtyRestocked   float32 `gorm:"TYPE:DECIMAL(12,4)"`
        QtyCanceled    float32 `gorm:"TYPE:DECIMAL(12,4)"`
        QtyBackordered float32 `gorm:"TYPE:DECIMAL(12,4)"`

        // Price
        Price               float64 `gorm:"TYPE:DECIMAL(20,6)"`
        PriceInclTax        float64 `gorm:"TYPE:DECIMAL(20,6)"`
        CustomPrice         float64 `gorm:"TYPE:DECIMAL(20,6)"`
        CustomPriceInclTax  float64 `gorm:"TYPE:DECIMAL(20,6)"`
        SpecialPrice        float64 `gorm:"TYPE:DECIMAL(20,6)"`
        SpecialPriceInclTax float64 `gorm:"TYPE:DECIMAL(20,6)"`

        // Discount
        DiscountAmount         float64 `gorm:"TYPE:DECIMAL(20,6)"`
        ShippingDiscountAmount float64 `gorm:"TYPE:DECIMAL(20,6)"`
        AppliedDiscounts       string  `gorm:"TYPE:VARCHAR(255)"`

        // Tax
        Taxable           bool
        TaxAmount         float64 `gorm:"TYPE:DECIMAL(20,6)"`
        DiscountTaxAmount float64 `gorm:"TYPE:DECIMAL(20,6)"`
        AppliedTaxes      string  `gorm:"TYPE:VARCHAR(255)"`

        // shipping
        ShippingAmount float64 `gorm:"TYPE:DECIMAL(20,6)"`

        // row total
        RowTotal        float64 `gorm:"TYPE:DECIMAL(20,6)"`
        RowTotalInclTax float64 `gorm:"TYPE:DECIMAL(20,6)"`

        Note     string `gorm:"TYPE:VARCHAR(255)"`
        MetaData string `gorm:"TYPE:VARCHAR(255)"`

        CreatedAt *time.Time
        UpdatedAt *time.Time
        DeletedAt *gorm.DeletedAt `gorm:"index"`
    }

    type SalesOrder struct {
        Id          string `gorm:"TYPE:VARCHAR(100);NOT NULL;PRIMARY_KEY"`
        OrderNumber string `gorm:"TYPE:VARCHAR(255)"`

        // status
        Status            string `gorm:"TYPE:VARCHAR(50)"`
        PaymentStatus     string `gorm:"TYPE:VARCHAR(50)"`
        FulfillmentStatus string `gorm:"TYPE:VARCHAR(50)"`
        ShipmentStatus    string `gorm:"TYPE:VARCHAR(50)"`
        ProcessStatus     string `gorm:"TYPE:VARCHAR(100)"`

        // total qty order_item
        TotalItemQty   float64 `gorm:"TYPE:DECIMAL(12,4)"`
        TotalItemCount int64

        // customer
        CustomerID          string `gorm:"TYPE:VARCHAR(100)"`
        CustomerExternalID  string `gorm:"TYPE:VARCHAR(100)"`
        CustomerEmail       string `gorm:"TYPE:VARCHAR(255)"`
        CustomerPhone       string `gorm:"TYPE:VARCHAR(255)"`
        CustomerFirstName   string `gorm:"TYPE:VARCHAR(255)"`
        CustomerMiddleName  string `gorm:"TYPE:VARCHAR(255)"`
        CustomerLastName    string `gorm:"TYPE:VARCHAR(255)"`
        CustomerTaxNumber   string `gorm:"TYPE:VARCHAR(255)"`
        CustomerCompanyName string `gorm:"TYPE:VARCHAR(255)"`

        // currency
        Currency             string  `gorm:"TYPE:VARCHAR(50)"`
        BaseCurrency         string  `gorm:"TYPE:VARCHAR(50)"`
        CurrencyExchangeRate float64 `gorm:"TYPE:DECIMAL(12,4)"`

        // subtotal
        Subtotal        float64 `gorm:"TYPE:DECIMAL(20,6)"`
        SubtotalInclTax float64 `gorm:"TYPE:DECIMAL(20,6)"`

        // discount
        DiscountPercent  float64 `gorm:"TYPE:DECIMAL(20,6)"`
        DiscountAmount   float64 `gorm:"TYPE:DECIMAL(20,6)"`
        AppliedDiscounts string  `gorm:"TYPE:VARCHAR(255)"`

        // shipping
        ShippingMethod      string  `gorm:"TYPE:VARCHAR(255)"`
        ShippingDescription string  `gorm:"TYPE:VARCHAR(255)"`
        ShippingAmount      float64 `gorm:"TYPE:DECIMAL(20,6)"`

        // tax
        TaxAmount float64 `gorm:"TYPE:DECIMAL(20,6)"`

        // grand total
        ExtraAmount float64 `gorm:"TYPE:DECIMAL(20,6)"`
        GrandTotal  float64 `gorm:"TYPE:DECIMAL(20,6)"`

        // meta data and additional
        Note     string `gorm:"TYPE:VARCHAR(255)"`
        Tags     string `gorm:"TYPE:VARCHAR(255)"`
        MetaData string `gorm:"TYPE:VARCHAR(255)"`

        CreatedAt *time.Time
        UpdatedAt *time.Time
        DeletedAt *gorm.DeletedAt `gorm:"index"`
    }

    return tx.AutoMigrate(
        &SalesOrder{},
        &SalesOrderItem{},
    )
}
