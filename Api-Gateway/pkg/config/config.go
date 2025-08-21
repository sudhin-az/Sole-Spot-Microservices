package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	UserSvcUrl    string `mapstructure:"USER_SVC_URL"`
	AddressSvcUrl string `mapstructure:"ADDRESS_SVC_URL"`
	AdminSvcUrl   string `mapstructure:"ADMIN_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	CartSvcUrl    string `mapstructure:"CART_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile("C:/Users/hp/Documents/Sole-Spot-Microservices/Api-Gateway/.env")
	viper.AddConfigPath("../..")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
