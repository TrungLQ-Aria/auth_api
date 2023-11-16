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
	UpdateRefreshToken(dx *gorm.DB, refresh *entity.AdminUser) error
	FindByRefreshToken(dx *gorm.DB, token string) (res *entity.AdminUser, err error)
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

func (*admin) UpdateRefreshToken(dx *gorm.DB, adminUser *entity.AdminUser) error {
	refresh := adminUser.Account.RefreshToken
	return dx.Model(refresh).
		Select("Token", "AccessTokenID", "issueAt").
		Updates(refresh).
		Error
}

func (*admin) FindByRefreshToken(dx *gorm.DB, token string) (res *entity.AdminUser, err error) {
	err = dx.
		Joins("JOIN accounts on accounts.id = admin_users.account_id").
		Joins("JOIN account_refresh_tokens on account_refresh_tokens.account_id = accounts.id").
		Preload("Account").
		Preload("Account.RefreshToken").
		Where("account_refresh_tokens.token = ?", token).
		First(&res).Error

	return res, err
}
