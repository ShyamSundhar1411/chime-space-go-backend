package bootstrap

import "github.com/ShyamSundhar1411/chime-space-go-backend/mongo"

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerPort             string `mapstructure:"SERVER_PORT"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBUser                 string `mapstructure:"DB_USER"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
}

type Application struct {
	Env   *Env
	Mongo mongo.Client
}
