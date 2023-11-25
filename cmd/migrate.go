package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "good_shoes/common/config"
    "good_shoes/logger"
    "good_shoes/migration"
)

var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "Migrate Command",
    Run:   runMirageCmd,
}

func init() {
    rootCmd.AddCommand(migrateCmd)
}

func runMirageCmd(cmd *cobra.Command, args []string) {
    config, err := config.LoadConfig(".")
    if err != nil {
        logger.Fatal("cannot load config:", err)
    }

    orm, err := gorm.Open(sqlite.Open(config.Database.Source), &gorm.Config{})
    if nil != err {
        fmt.Println(err)
        os.Exit(1)
    }

    sqliteDB, err := orm.DB()
    if nil != err {
        panic(err)
    }

    defer func() {
        if err := sqliteDB.Close(); err != nil {
            panic(err)
        }
    }()

    fmt.Println("mysql connection established")
    err = migration.Migrate(orm)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("migration successful")
    os.Exit(0)
}
