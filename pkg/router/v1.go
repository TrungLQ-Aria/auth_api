package router

import (
	"ema_sound_clone_api/config"
	"ema_sound_clone_api/pkg/handler"
	routermiddleware "ema_sound_clone_api/pkg/router/middleware"
	"github.com/labstack/echo/v4"
)

func VersionOne(v1 *echo.Group, env config.Env) {
	adminGroup := v1.Group("/admins")

	var (
		h = handler.NewAdmin()
		a = routermiddleware.NewAuthMiddleware(env)
	)

	adminGroup.POST("", h.CreateAdminUserByDev, a.DevAPIKeyAuthentication())

	adminGroup.POST("/sign-in", h.SignIn)

	adminGroup.POST("/access-token", h.RefreshToken)

	adminGroup.GET("", h.GetAllAdminUser, a.AdminAuthentication())
}
