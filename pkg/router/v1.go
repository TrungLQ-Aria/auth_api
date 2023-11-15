package router

import (
	"ema_sound_clone_api/pkg/handler"
	"github.com/labstack/echo/v4"
)

func VersionOne(v1 *echo.Group) {
	g := v1.Group("/admins")

	var (
		h = handler.NewAdmin()
	)

	g.POST("", h.CreateAdminUserByDev)
}
