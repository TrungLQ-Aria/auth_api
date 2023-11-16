package router

import (
	"ema_sound_clone_api/config"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, env config.Env) {
	v1 := e.Group("/v1")

	VersionOne(v1, env)
}
