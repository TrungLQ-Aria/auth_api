package server

import (
	"ema_sound_clone_api/config"
	"ema_sound_clone_api/internal/db"
	"ema_sound_clone_api/pkg/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
