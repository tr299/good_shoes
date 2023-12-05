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
    BaseUrl       string
    ApiPrefix     string
    LoggerConfig  LoggerConfig
    Tracer        TracerConfig
    Database      *DbConfig
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
    config.ApiPrefix = viper.GetString("server.api_prefix")

    config.Database = &DbConfig{
        Driver: "database.driver",
        Source: "database.source",
    }

    return
}
