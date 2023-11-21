package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/trungaria/auth_api.git/config"
	"github.com/trungaria/auth_api.git/internal/db"
	"github.com/trungaria/auth_api.git/pkg/router"
)

func Boostrap(e *echo.Echo) {
	env, err := config.NewEnv()
	if err != nil {
		log.Panicf("Failed create env", err)
	}

	db.Connect(*env)

	if err = db.Migrate(); err != nil {
		log.Panicf("Failed migrate")
	}

	router.Init(e, *env)
}
