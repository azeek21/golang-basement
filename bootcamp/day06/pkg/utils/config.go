package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig(stage string) error {
	godotenv.Load(".env." + stage)
	viper.AddConfigPath(fmt.Sprintf(`./config/%v/`, stage))

	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func LoadConfig[T interface{}](config T) (T, error) {
	err := viper.Unmarshal(&config)
	return config, err
}

func Must(err error) {
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
}
