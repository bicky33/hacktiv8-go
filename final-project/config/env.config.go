package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type ConfigApp struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_DATABASE_NAME"`
	DbSslMode  string `mapstructure:"DB_SSL_MODE"`
	DbPort     int    `mapstructure:"DB_PORT"`

	AppPort int `mapstructure:"APP_PORT"`

	AccessTokenPrivateKey string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey  string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge     int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
}

func Config() (config ConfigApp) {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigType("env")
	vp.SetConfigName(".env")
	vp.AutomaticEnv()

	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf("Read error %v", err)
	}

	if err := vp.Unmarshal(&config); err != nil {
		log.Fatalln("unable to unmarshall the config ", err)
	}
	return
}
