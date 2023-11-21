package main

import (
	"github.com/trungaria/auth_api.git/config"
	"github.com/trungaria/auth_api.git/pkg/server"
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
