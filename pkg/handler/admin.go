package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/trungaria/auth_api.git/internal/utils/response"
	"github.com/trungaria/auth_api.git/pkg/handler/openapi"
	"github.com/trungaria/auth_api.git/pkg/query"
	"github.com/trungaria/auth_api.git/pkg/usecase"
)

func (h OpenAPIHandler) PostV1Admins(c echo.Context) error {
	var (
		req openapi.AdminUserCreateRequest
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

func (h OpenAPIHandler) PostV1AdminUserSignIn(c echo.Context) error {
	var (
		req openapi.AdminUserSignInRequest
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

func (h OpenAPIHandler) PostV1AdminUserAccessToken(c echo.Context, params openapi.PostV1AdminUserAccessTokenParams) error {
	res, err := h.AdminUsecase.RefreshToken(params)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res)
}

func (h OpenAPIHandler) GetV1AdminUsers(c echo.Context) error {
	var (
		q = query.NewAdmin()
	)

	res, err := q.FindAllAdminUser()
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res)
}
