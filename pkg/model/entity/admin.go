package entity

import "gorm.io/gorm"

type (
	AdminUserID uint
	AdminUser   struct {
		gorm.Model
		AccountID uint     `gorm:"notNull; unique;"`
		Account   *Account ``
		Email     string   `gorm:"size:256; notNull; unique;"`
		Name      string   `gorm:"size:64; notNull;"`
	}
)
