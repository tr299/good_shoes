package versions

import (
    "gorm.io/gorm"
    "time"
)

func Version20231116000000(tx *gorm.DB) error {
    type Category struct {
        Id          string `gorm:"TYPE:VARCHAR(100);NOT NULL;PRIMARY_KEY"`
        ParentID    string `gorm:"TYPE:VARCHAR(100)"`
        Name        string `gorm:"TYPE:VARCHAR(255)"`
        Description string `gorm:"TYPE:LONGTEXT"`
        Status      string `gorm:"TYPE:VARCHAR(100)"`
        Path        string `gorm:"TYPE:VARCHAR(500)"`
        CreatedAt   *time.Time
        UpdatedAt   *time.Time
        DeletedAt   *gorm.DeletedAt `gorm:"index"`
    }

    type OptionItem struct {
        Id         string  `gorm:"TYPE:VARCHAR(100);NOT NULL;PRIMARY_KEY"`
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
        CreatedAt  *time.Time
        UpdatedAt  *time.Time
        DeletedAt  *gorm.DeletedAt `gorm:"index"`
    }

    type Option struct {
        Id         string  `gorm:"TYPE:VARCHAR(100);NOT NULL;PRIMARY_KEY"`
        ProductID  string  `gorm:"TYPE:VARCHAR(255);"`
        Key        string  `gorm:"TYPE:VARCHAR(255);"`
        Name       string  `gorm:"TYPE:VARCHAR(255);"`
        Note       string  `gorm:"TYPE:VARCHAR(255);"`
        Position   uint32  `gorm:"TYPE:INT;"`
        IsRequired bool    `gorm:"TYPE:TINYINT(1);NOT NULL;default:0;"`
        Type       string  `gorm:"TYPE:VARCHAR(50);"`
        Qty        float32 `gorm:"TYPE:DECIMAL(12, 4);"`
        Price      float64 `gorm:"TYPE:DECIMAL(20,4)"`
        PriceType  string  `gorm:"TYPE:VARCHAR(30);"`
        CreatedAt  *time.Time
        UpdatedAt  *time.Time
        DeletedAt  *gorm.DeletedAt `gorm:"index"`
    }

    type Product struct {
        Id               string  `gorm:"TYPE:VARCHAR(100);NOT NULL;PRIMARY_KEY"`
        IsVariant        bool    `gorm:"index:platform_external_idx,priority:2;TYPE:TINYINT(1);NOT NULL;default:0;"`
        Sku              string  `gorm:"TYPE:VARCHAR(255);"`
        Name             string  `gorm:"TYPE:VARCHAR(255);"`
        Description      string  `gorm:"TYPE:LONGTEXT;"`
        Description2     string  `gorm:"TYPE:LONGTEXT;"`
        Status           string  `gorm:"TYPE:VARCHAR(20);"`
        Barcode          string  `gorm:"TYPE:VARCHAR(255);"`
        Type             string  `gorm:"TYPE:VARCHAR(255);"`
        OptionKey        string  `gorm:"TYPE:VARCHAR(255);"`
        OptionValue      string  `gorm:"TYPE:VARCHAR(255);"`
        Position         string  `gorm:"TYPE:VARCHAR(255);"`
        Tags             string  `gorm:"TYPE:VARCHAR(255);"`
        Price            float64 `gorm:"TYPE:DECIMAL(20,4)"`
        SalePrice        float64 `gorm:"TYPE:DECIMAL(20,4)"`
        Cost             float64 `gorm:"TYPE:DECIMAL(20,4)"`
        CategoryIds      string  `gorm:"TYPE:TEXT"`
        CheckInventory   bool
        MultipleVariants bool
        TotalQuantity    uint32 `gorm:"TYPE:INT;"`
        Brand            string `gorm:"TYPE:VARCHAR(100);"`
        ImageUrl         string `gorm:"TYPE:LONGTEXT;"`
        CreatedAt        *time.Time
        UpdatedAt        *time.Time
        DeletedAt        *gorm.DeletedAt `gorm:"index"`
    }

    type ProductCategory struct {
        Id         string `gorm:"TYPE:VARCHAR(100);NOT NULL;PRIMARY_KEY"`
        ProductID  string `gorm:"uniqueIndex:idx_product_category;uniqueIndex:platform_category_external_idx;TYPE:VARCHAR(255)"`
        CategoryID string `gorm:"uniqueIndex:idx_product_category;TYPE:VARCHAR(255)"`
        CreatedAt  *time.Time
        UpdatedAt  *time.Time
        DeletedAt  *gorm.DeletedAt `gorm:"index"`
    }

    return tx.AutoMigrate(
        &Option{},
        &OptionItem{},
        &Product{},
        &Category{},
    )
}
