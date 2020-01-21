package config

import (
	"fmt"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func Load() {
	viper.SetDefault("port", 3000)
	viper.SetConfigType("yaml")
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func AviatorBaseURL() *url.URL {
	aviatorBaseURL := GetValue("AVIATOR_BASE_URL")
	baseURL, _ := url.Parse(aviatorBaseURL)
	return baseURL
}

func AviatorApiKey() string {
	return GetValue("AVIATOR_API_KEY")
}
