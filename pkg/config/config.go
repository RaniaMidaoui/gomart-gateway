package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SERVICE_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SERVICE_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SERVICE_URL"`
}

func SetDefaultConfig() {
	viper.SetDefault("PORT", ":3000")
	viper.SetDefault("AUTH_SERVICE_URL", "127.0.0.1:50051")
	viper.SetDefault("PRODUCT_SERVICE_URL", "127.0.0.1:50052")
	viper.SetDefault("ORDER_SERVICE_URL", "127.0.0.1:50053")
}

func ReadSystemEnv() {

	env_port, ok := os.LookupEnv("PORT")
	if !ok || env_port == "" {
		fmt.Println("System environment variables 'PORT' not set, Working with default values instead...")
		viper.SetDefault("PORT", ":3000")
	}
	viper.BindEnv("PORT")

	env_auth, ok := os.LookupEnv("AUTH_SERVICE_URL")
	if !ok || env_auth == "" {
		fmt.Println("System environment variables 'AUTH_SERVICE_URL' not set, Working with default values instead...")
		viper.SetDefault("AUTH_SERVICE_URL", "127.0.0.1:50051")
	}
	viper.BindEnv("AUTH_SERVICE_URL")

	env_product, ok := os.LookupEnv("PRODUCT_SERVICE_URL")
	if !ok || env_product == "" {
		fmt.Println("System environment variables 'PRODUCT_SERVICE_URL' not set, Working with default values instead...")
		viper.SetDefault("PRODUCT_SERVICE_URL", "127.0.0.1:50052")
	}
	viper.BindEnv("PRODUCT_SERVICE_URL")

	env_order, ok := os.LookupEnv("ORDER_SERVICE_URL")
	if !ok || env_order == "" {
		fmt.Println("System environment variables 'ORDER_SERVICE_URL' not set, Working with default values instead...")
		viper.SetDefault("ORDER_SERVICE_URL", "127.0.0.1:50053")
	}
	viper.BindEnv("ORDER_SERVICE_URL")
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No configuration file found, checking system environment variables instead ...")
			ReadSystemEnv()
		} else {
			fmt.Println("Error proceeded while trying to set environment variables, check your configuration")
			fmt.Println("Working with default values instead...")
			SetDefaultConfig()
		}
	}

	err = viper.Unmarshal(&c)
	fmt.Println(os.LookupEnv("DB_URL"))
	fmt.Println(c)
	return
}
