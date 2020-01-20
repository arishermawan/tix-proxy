package config

import (
	"github.com/spf13/viper"
	"os"
)

func GetValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}
