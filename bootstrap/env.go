package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Unable to load environment file :", err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Unable to decode into struct :", err)
	}
	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
	return &env
}
