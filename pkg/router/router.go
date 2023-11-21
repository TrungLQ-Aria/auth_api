package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trungaria/auth_api.git/config"
	"github.com/trungaria/auth_api.git/pkg/handler"
	"github.com/trungaria/auth_api.git/pkg/query"
	"github.com/trungaria/auth_api.git/pkg/usecase"
)

func Init(e *echo.Echo, env config.Env) {
	v1 := e.Group("/v1")

	h := handler.OpenAPIHandler{
		AdminUsecase: usecase.NewAdmin(),
		AdminQuery:   query.NewAdmin(),
	}
	VersionOne(v1, env, h)
}
