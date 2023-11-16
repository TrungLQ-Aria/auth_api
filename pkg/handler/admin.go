package handler

import (
	"ema_sound_clone_api/internal/utils/response"
	"ema_sound_clone_api/pkg/model/request"
	"ema_sound_clone_api/pkg/query"
	"ema_sound_clone_api/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type Admin interface {
	CreateAdminUserByDev(c echo.Context) error
	SignIn(c echo.Context) error
	RefreshToken(c echo.Context) error
	GetAllAdminUser(c echo.Context) error
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}

func (admin) CreateAdminUserByDev(c echo.Context) error {
	var (
		req request.AdminUserCreateByDevRequest
		uc  = usecase.NewAdmin()
	)

	if err := c.Bind(&req); err != nil {
		return response.R400(c, nil, err.Error())
	}

	res, err := uc.CreateUserAdminByDev(req)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res)
}

func (admin) SignIn(c echo.Context) error {
	var (
		req request.AdminUserLoginRequest
		uc  = usecase.NewAdmin()
	)

	if err := c.Bind(&req); err != nil {
		return response.R400(c, nil, err.Error())
	}

	res, err := uc.SignIn(req)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res)
}

func (admin) RefreshToken(c echo.Context) error {
	var (
		req request.AdminUserRefreshToken
		uc  = usecase.NewAdmin()
	)

	if err := c.Bind(&req); err != nil {
		return response.R400(c, nil, err.Error())
	}

	res, err := uc.RefreshToken(req)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res)
}

func (admin) GetAllAdminUser(c echo.Context) error {
	var (
		q = query.NewAdmin()
	)

	res, err := q.FindAllAdminUser()
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res)
}
