package main

import (
	"ema_sound_clone_api/config"
	"ema_sound_clone_api/pkg/server"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	server.Boostrap(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.GetEnv().AppPort)))
}
