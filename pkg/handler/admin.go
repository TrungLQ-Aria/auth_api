package handler

import (
	"ema_sound_clone_api/internal/utils/response"
	"github.com/labstack/echo/v4"
)

type Admin interface {
	CreateAdminUserByDev(c echo.Context) error
	SignIn(c echo.Context) error
	RefreshToken(c echo.Context) error
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}

func (admin) CreateAdminUserByDev(c echo.Context) error {
	return response.R200(c, nil)
}

func (admin) SignIn(c echo.Context) error {
	return response.R200(c, nil)
}

func (admin) RefreshToken(c echo.Context) error {
	return response.R200(c, nil)
}
