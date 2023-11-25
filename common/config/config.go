package config

import (
    "github.com/spf13/viper"
)

type LoggerConfig struct {
    Level       string
    JSONFormat  bool
    EnableTrace bool
}

type TracerConfig struct {
    Endpoint       string
    Username       string
    Password       string
    ServiceName    string
    ServiceVersion string
    Environment    string
}

type DbConfig struct {
    Driver string
    Source string
}

type Config struct {
    ServerAddress string
    LoggerConfig  LoggerConfig
    Tracer        TracerConfig
    Database      *DbConfig
    ProductConfig *ProductConfig
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }
    config.ServerAddress = viper.GetString("server.address")

    config.Database = &DbConfig{
        Driver: "database.driver",
        Source: "database.source",
    }

    config.ProductConfig = &ProductConfig{
        BaseUrl:   viper.GetString("product.base_url"),
        ApiPrefix: viper.GetString("product.api_prefix"),
    }

    return
}
