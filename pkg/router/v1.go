package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trungaria/auth_api.git/config"
	"github.com/trungaria/auth_api.git/pkg/handler/openapi"
	routermiddleware "github.com/trungaria/auth_api.git/pkg/router/middleware"
)

func VersionOne(v1 *echo.Group, env config.Env, si openapi.ServerInterface) {
	adminGroup := v1.Group("/admins")

	var (
		authDev   = routermiddleware.DevAPIKeyAuthentication(env)
		authAdmin = routermiddleware.AdminAuthentication(env)
	)

	wrap := openapi.ServerInterfaceWrapper{
		Handler: si,
	}

	adminGroup.POST("", wrap.PostV1Admins, authDev)

	adminGroup.POST("/sign-in", wrap.PostV1AdminUserSignIn)

	adminGroup.POST("/access-token", wrap.PostV1AdminUserAccessToken)

	adminGroup.GET("", wrap.GetV1AdminUsers, authAdmin)
}
