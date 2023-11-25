package migration

import (
    "github.com/go-gormigrate/gormigrate/v2"
    "good_shoes/migration/versions"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID:      "20231116000000",
            Migrate: versions.Version20231116000000,
        },
    })

    return m.Migrate()
}
