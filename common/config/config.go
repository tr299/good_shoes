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
	WspGoConfig   *WspGoConfig
	MspGoConfig   *MspGoConfig
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

	config.WspGoConfig = &WspGoConfig{
		BaseUrl:   viper.GetString("wsp-go.base_url"),
		ApiPrefix: viper.GetString("wsp-go.api_prefix"),
	}

	config.MspGoConfig = &MspGoConfig{
		BaseUrl: viper.GetString("msp-go.base_url"),
	}

	return
}
