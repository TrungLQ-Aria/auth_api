package query

import (
	"github.com/trungaria/auth_api.git/internal/db"
	"github.com/trungaria/auth_api.git/pkg/handler/openapi"
	"github.com/trungaria/auth_api.git/pkg/model/entity"
)

type Admin interface {
	FindAllAdminUser() (*openapi.AdminUserResponse, error)
	ConvertToResponse(adminUser entity.AdminUser) openapi.AdminUser
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}

func (*admin) ConvertToResponse(adminUser entity.AdminUser) openapi.AdminUser {
	return openapi.AdminUser{
		Name:  adminUser.Name,
		Email: adminUser.Email,
	}
}

func (q *admin) FindAllAdminUser() (res *openapi.AdminUserResponse, err error) {
	var (
		d      = db.GetDb()
		founds = []entity.AdminUser{}
	)

	err = d.Model(&entity.AdminUser{}).Find(&founds).Error

	if err != nil {
		return nil, err
	}

	adminUsers := make([]openapi.AdminUser, len(founds))

	for i, f := range founds {
		adminUsers[i] = q.ConvertToResponse(f)
	}

	return &openapi.AdminUserResponse{
		AdminUsers: &adminUsers,
	}, nil
}
