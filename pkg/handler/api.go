package handler

import (
	"github.com/trungaria/auth_api.git/pkg/query"
	"github.com/trungaria/auth_api.git/pkg/usecase"
)

type OpenAPIHandler struct {
	AdminUsecase usecase.Admin
	AdminQuery   query.Admin
}
