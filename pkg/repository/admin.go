package repository

import (
	"ema_sound_clone_api/internal/db"
	"ema_sound_clone_api/pkg/model/entity"
	"gorm.io/gorm"
)

type Admin interface {
	FindBySignIn(dx *gorm.DB, signIn string) (res *entity.AdminUser, err error)
	FindBy(dx *gorm.DB, predicate *entity.AdminUser) (res *entity.AdminUser, err error)
	Create(dx *gorm.DB, adminUser *entity.AdminUser) error
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}

func (r *admin) FindBySignIn(dx *gorm.DB, signIn string) (res *entity.AdminUser, err error) {
	if dx == nil {
		dx = db.GetDb()
	}

	return r.FindBy(dx, &entity.AdminUser{Email: signIn})
}

func (*admin) FindBy(dx *gorm.DB, predicate *entity.AdminUser) (res *entity.AdminUser, err error) {
	if dx == nil {
		dx = db.GetDb()
	}

	err = dx.Preload("Account").
		Preload("Account.RefreshToken").
		Where(&predicate).
		First(&res).
		Error

	return
}

func (*admin) Create(dx *gorm.DB, adminUser *entity.AdminUser) error {
	return dx.Create(adminUser).Error
}
