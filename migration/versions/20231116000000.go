package versions

import "gorm.io/gorm"

func Version20231116000000(tx *gorm.DB) error {
    type Category struct {
        Id          string `gorm:"TYPE:VARCHAR(100)"`
        ParentID    string `gorm:"TYPE:VARCHAR(100)"`
        Name        string `gorm:"TYPE:VARCHAR(255)"`
        Description string `gorm:"TYPE:LONGTEXT"`
        Status      string `gorm:"TYPE:VARCHAR(100)"`
        Path        string `gorm:"TYPE:VARCHAR(500)"`
    }

    type OptionItem struct {
        Id         string  `gorm:"TYPE:VARCHAR(100)"`
        OptionID   string  `gorm:"TYPE:VARCHAR(255);"`
        Label      string  `gorm:"TYPE:VARCHAR(255);"`
        Value      string  `gorm:"TYPE:VARCHAR(255);"`
        Price      float64 `gorm:"TYPE:DECIMAL(20,4)"`
        PriceType  string  `gorm:"TYPE:VARCHAR(30);"`
        Sku        string  `gorm:"TYPE:VARCHAR(255);"`
        Qty        float32 `gorm:"TYPE:DECIMAL(12, 4);"`
        Position   uint32  `gorm:"TYPE:INT;"`
        IsChecked  bool    `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        IsSelected bool    `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        IsDefault  bool    `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        IsDisabled bool    `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        QtyMutable bool    `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        ProductSku string  `gorm:"TYPE:TEXT;"`
    }

    type Option struct {
        Id         string       `gorm:"TYPE:VARCHAR(100)"`
        ProductID  string       `gorm:"TYPE:VARCHAR(255);"`
        Key        string       `gorm:"TYPE:VARCHAR(255);"`
        Name       string       `gorm:"TYPE:VARCHAR(255);"`
        Note       string       `gorm:"TYPE:VARCHAR(255);"`
        Position   uint32       `gorm:"TYPE:INT;"`
        IsRequired bool         `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        Type       string       `gorm:"TYPE:VARCHAR(50);"`
        Qty        float32      `gorm:"TYPE:DECIMAL(12, 4);"`
        Price      float64      `gorm:"TYPE:DECIMAL(20,4)"`
        PriceType  string       `gorm:"TYPE:VARCHAR(30);"`
        Items      []OptionItem `gorm:"foreignKey:OptionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    }

    type Product struct {
        Id                          string `gorm:"TYPE:VARCHAR(100)"`
        StoreID                     string `gorm:"index:platform_external_idx,priority:1;TYPE:VARCHAR(100);"`
        IsVariant                   bool   `gorm:"index:platform_external_idx,priority:2;TYPE:TINYINT(1);NOT NULL;default:0;"`
        InventoryItemID             string `gorm:"TYPE:VARCHAR(100);"`
        ParentID                    string `gorm:"TYPE:VARCHAR(100);"`
        Sku                         string `gorm:"TYPE:VARCHAR(255);"`
        Name                        string `gorm:"TYPE:VARCHAR(255);"`
        Handle                      string `gorm:"TYPE:TEXT;"`
        ImageID                     string `gorm:"TYPE:VARCHAR(255);"`
        Description                 string `gorm:"TYPE:LONGTEXT;"`
        Status                      string `gorm:"TYPE:VARCHAR(20);"`
        Visibility                  string `gorm:"TYPE:VARCHAR(20);"`
        Barcode                     string `gorm:"TYPE:VARCHAR(255);"`
        Manufacturer                string `gorm:"TYPE:VARCHAR(255);"`
        Gtn                         string `gorm:"TYPE:VARCHAR(255);"`
        Gtin                        string `gorm:"TYPE:VARCHAR(255);"`
        Ean                         string `gorm:"TYPE:VARCHAR(255);"`
        Upc                         string `gorm:"TYPE:VARCHAR(255);"`
        Mpn                         string `gorm:"TYPE:VARCHAR(255);"`
        Bpn                         string `gorm:"TYPE:VARCHAR(255);"`
        Type                        string `gorm:"TYPE:VARCHAR(255);"`
        TaxId                       string `gorm:"TYPE:VARCHAR(255);"`
        Weight                      string `gorm:"TYPE:VARCHAR(20);"`
        Width                       string `gorm:"TYPE:VARCHAR(20);"`
        Height                      string `gorm:"TYPE:VARCHAR(20);"`
        Length                      string `gorm:"TYPE:VARCHAR(20);"`
        Process                     string `gorm:"TYPE:VARCHAR(255);"`
        OptionKey                   string `gorm:"TYPE:VARCHAR(255);"`
        OptionValue                 string `gorm:"TYPE:VARCHAR(255);"`
        Position                    string `gorm:"TYPE:VARCHAR(255);"`
        IsVirtual                   bool   `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        RequireShipping             bool   `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        AllowShipping               bool   `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        FreeShipping                bool   `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        Taxable                     bool   `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        Tags                        string `gorm:"TYPE:VARCHAR(255);"`
        Msrp                        string `gorm:"TYPE:varchar(50);"`
        Price                       string `gorm:"TYPE:varchar(50);NOT NULL;default:0;"`
        CompareAtPrice              string `gorm:"TYPE:varchar(50);"`
        Cost                        string `gorm:"TYPE:varchar(50);"`
        SpecialPrice                string `gorm:"TYPE:varchar(50);"`
        SpecialPriceFrom            string `gorm:"TYPE:varchar(50);"`
        SpecialPriceTo              string `gorm:"TYPE:varchar(50);"`
        WarehouseIds                string `gorm:"TYPE:TEXT"`
        CategoryIds                 string `gorm:"TYPE:TEXT"`
        Backorders                  bool
        CheckInventory              bool
        MultipleVariants            bool
        TotalAvailableWarehouseItem string     `gorm:"TYPE:VARCHAR(100);"`
        Categories                  []Category `gorm:"many2many:product_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
        Options                     []Option   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    }

    type ProductCategory struct {
        Id         string `gorm:"TYPE:VARCHAR(100)"`
        ProductID  string `gorm:"uniqueIndex:idx_product_category;uniqueIndex:platform_category_external_idx;TYPE:VARCHAR(255)"`
        CategoryID string `gorm:"uniqueIndex:idx_product_category;TYPE:VARCHAR(255)"`
        StoreID    string `gorm:"uniqueIndex:platform_category_external_idx;TYPE:VARCHAR(100);"`
        ExternalID string `gorm:"uniqueIndex:platform_category_external_idx;TYPE:VARCHAR(100);"`
    }

    err := tx.SetupJoinTable(&Product{}, "Categories", &ProductCategory{})
    if err != nil {
        return err
    }

    return tx.AutoMigrate(
        &Option{},
        &OptionItem{},
        &Product{},
        &Category{},
    )
}
