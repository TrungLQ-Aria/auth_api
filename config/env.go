package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	AppEnv  string `envconfig:"APP_ENV"`
	AppPort string `envconfig:"APP_PORT"`

	// Mysql
	MysqlUser     string `envconfig:"MYSQL_USER"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD"`
	MysqlDatabase string `envconfig:"MYSQL_DATABASE"`
	MysqlProtocol string `envconfig:"MYSQL_PROTOCOL"`

	// APIKEY
	DevApiKey string `envconfig:"DEV_API_KEY"`

	// JWT KEY
	AdminJWTKey string `envconfig:"ADMIN_JWT_SECRET_KEY"`
}

var (
	env Env
)

func NewEnv() (*Env, error) {
	if err := envconfig.Process("", &env); err != nil {
		return nil, err
	}

	fmt.Println(env)

	return &env, nil
}

func GetEnv() Env {
	return env
}
