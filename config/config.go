package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Database DatabaseConfig
	AuthKey  KeyAuth
	Server   Server
	Auth     Auth
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SslMode  string
	Schema   Schema
}

type Schema struct {
	App string
}

type KeyAuth struct {
	PaSecret string
	DaSecret string
}

type Server struct {
	Port         string
	Issuer       string
	Env          string
	BaseUrl      string
	RedirectPage string
}

type Auth struct {
	CmsSecret         string
	ApplicationIssuer string
	ExpiredTimeHour   int
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln(err)
	}
}
