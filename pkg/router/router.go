package router

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	v1 := e.Group("/v1")

	VersionOne(v1)
}
