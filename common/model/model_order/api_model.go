package model_order

type SalesOrderItem struct {
    Id              string `json:"id"`
    SalesOrderID    string `json:"sales_order_id"`
    ProductID       string `json:"product_id"`
    ParentProductID string `json:"parent_product_id"`
    Type            string `json:"type"`
    Sku             string `json:"sku"`
    Name            string `json:"name"`
    MediaUrl        string `json:"mediaUrl"`
    IsVirtual       bool   `json:"is_virtual"`
    IsGiftProduct   bool   `json:"is_gift_product"`
    RequireShipping bool   `json:"require_shipping"`

    // Quantity
    QtyOrdered     float32 `json:"qty_ordered"`
    QtyInvoiced    float32 `json:"qty_invoiced"`
    QtyFulfilled   float32 `json:"qty_fulfilled"`
    QtyShipped     float32 `json:"qty_shipped"`
    QtyRefunded    float32 `json:"qty_refunded"`
    QtyRestocked   float32 `json:"qty_restocked"`
    QtyCanceled    float32 `json:"qty_canceled"`
    QtyBackordered float32 `json:"qty_backordered"`

    // Price
    Price               float64 `json:"price"`
    PriceInclTax        float64 `json:"price_incl_tax"`
    CustomPrice         float64 `json:"sale_price"`
    CustomPriceInclTax  float64 `json:"custom_price_incl_tax"`
    SpecialPrice        float64 `json:"special_price"`
    SpecialPriceInclTax float64 `json:"special_price_incl_tax"`

    // Discount
    DiscountAmount         float64 `json:"discount_amount"`
    ShippingDiscountAmount float64 `json:"shipping_discount_amount"`
    AppliedDiscounts       string  `json:"applied_discounts"`

    // Tax
    Taxable           bool    `json:"taxable"`
    TaxAmount         float64 `json:"tax_amount"`
    DiscountTaxAmount float64 `json:"discount_tax_amount"`
    AppliedTaxes      string  `json:"applied_taxes"`

    // shipping
    ShippingAmount float64 `json:"shipping_amount"`

    // row total
    RowTotal        float64 `json:"row_total"`
    RowTotalInclTax float64 `json:"row_total_incl_tax"`

    Note     string `json:"note"`
    MetaData string `json:"meta_data"`

    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt string `json:"deleted_at,omitempty"`
}

type SalesOrder struct {
    Id          string `json:"id"`
    OrderNumber string `json:"order_number"`

    // status
    Status            string `json:"status"`
    PaymentStatus     string `json:"payment_status"`
    FulfillmentStatus string `json:"fulfillment_status"`
    ShipmentStatus    string `json:"shipment_status"`
    ProcessStatus     string `json:"process_status"`

    // total qty order_item
    TotalItemQty   float64 `json:"total_item_qty"`
    TotalItemCount int64   `json:"total_item_count"`

    // customer
    CustomerID          string `json:"customer_id"`
    CustomerExternalID  string `json:"customer_external_id"`
    CustomerEmail       string `json:"customer_email"`
    CustomerPhone       string `json:"customer_phone"`
    CustomerFirstName   string `json:"customer_first_name"`
    CustomerMiddleName  string `json:"customer_middle_name"`
    CustomerLastName    string `json:"customer_last_name"`
    CustomerTaxNumber   string `json:"customer_tax_number"`
    CustomerCompanyName string `json:"customer_company_name"`

    // currency
    Currency             string  `json:"currency"`
    BaseCurrency         string  `json:"base_currency"`
    CurrencyExchangeRate float64 `json:"currency_exchange_rate"`

    // subtotal
    Subtotal        float64 `json:"subtotal"`
    SubtotalInclTax float64 `json:"subtotal_incl_tax"`

    // discount
    DiscountPercent  float64 `json:"discount_percent"`
    DiscountAmount   float64 `json:"discount_amount"`
    AppliedDiscounts string  `json:"applied_discounts"`

    // shipping
    ShippingMethod      string  `json:"shipping_method"`
    ShippingDescription string  `json:"shipping_description"`
    ShippingAmount      float64 `json:"shipping_amount"`

    // tax
    TaxAmount float64 `json:"tax_amount"`

    // grand total
    ExtraAmount float64 `json:"extra_amount"`
    GrandTotal  float64 `json:"grand_total"`

    // meta data and additional
    Note     string `json:"note"`
    Tags     string `json:"tags"`
    MetaData string `json:"meta_data"`

    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt string `json:"deleted_at,omitempty"`
}

type SalesOrderDetail struct {
    SalesOrder
    Items []SalesOrderItem `json:"items,omitempty"`
}

type ListSalesOrderResponse struct {
    Total int64        `json:"total"`
    Items []SalesOrder `json:"items"`
}

type ListSalesOrderRequest struct {
    Page   int    `form:"page"`
    Limit  int    `form:"limit"`
    Search string `form:"search"`
}

type GetOrderByIdRequest struct {
    Id string `uri:"id"`
}

type CreateSalesOrderRequest struct {
    SalesOrderDetail
}

type CreateSalesOrderResponse struct {
    Message string `json:"message"`
    OrderId string `json:"order_id"`
}

type UpdateOrderStatusRequest struct {
    Id     string `uri:"id"`
    Status string `json:"status"`
}

type UpdateOrderStatusResponse struct {
    OrderId string `json:"order_id"`
    Status  string `json:"status"`
    Message string `json:"message"`
}
