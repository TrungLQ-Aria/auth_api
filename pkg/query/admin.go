package query

import (
	"ema_sound_clone_api/internal/db"
	"ema_sound_clone_api/pkg/model/entity"
	responsemodel "ema_sound_clone_api/pkg/model/response"
)

type Admin interface {
	FindAllAdminUser() (*responsemodel.ListAdminUsersResponse, error)
	ConvertToResponse(adminUser entity.AdminUser) responsemodel.AdminUserDetail
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}

func (*admin) ConvertToResponse(adminUser entity.AdminUser) responsemodel.AdminUserDetail {
	return responsemodel.AdminUserDetail{
		Name:      adminUser.Name,
		Email:     adminUser.Email,
		CreatedAt: adminUser.CreatedAt,
		UpdatedAt: adminUser.UpdatedAt,
	}
}

func (q *admin) FindAllAdminUser() (res *responsemodel.ListAdminUsersResponse, err error) {
	var (
		d      = db.GetDb()
		founds = []entity.AdminUser{}
	)

	err = d.Model(&entity.AdminUser{}).Find(&founds).Error

	if err != nil {
		return nil, err
	}

	adminUsers := make([]responsemodel.AdminUserDetail, len(founds))

	for i, f := range founds {
		adminUsers[i] = q.ConvertToResponse(f)
	}

	return &responsemodel.ListAdminUsersResponse{
		AdminUsers: adminUsers,
	}, nil
}
