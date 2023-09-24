package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppConfig struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

func BuildConfig(l *zap.Logger) *AppConfig {
	viper.SetConfigName("app-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		l.Error("Error while reading configuration file",
			zap.String("err", err.Error()))
	}

	var config AppConfig

	if err := viper.Unmarshal(&config); err != nil {
		l.Error("Error unmarshaling configuration",
			zap.String("err", err.Error()))
	}

	return &config
}
