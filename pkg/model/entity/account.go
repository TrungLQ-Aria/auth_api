package entity

import (
	"gorm.io/gorm"
	"time"
)

type (
	AccountID uint
	Account   struct {
		gorm.Model
		Password     string               `gorm:"size:256; notNull"`
		SignedUpAt   time.Time            `gorm:"notNull;"`
		AdminUser    *AdminUser           ``
		RefreshToken *AccountRefreshToken ``
	}
	AccountRefreshToken struct {
		gorm.Model
		AccountID     uint      `gorm:"notNull; unique;"`
		Account       *Account  ``
		Token         string    `gorm:"size:36; notNull; unique;"`
		AccessTokenID string    `gorm:"size:36; notNull;"`
		IssuedAt      time.Time `gorm:"notNull"`
	}
)

const (
	refreshTokenExpiration = 24 * time.Hour * 30
)

func (e *AccountRefreshToken) ExpiredAt() time.Time {
	return e.IssuedAt.Add(refreshTokenExpiration)
}

func (e *AccountRefreshToken) Expired(now time.Time) bool {
	return !e.ExpiredAt().After(now)
}
