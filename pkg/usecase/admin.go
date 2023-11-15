package usecase

import (
	db2 "ema_sound_clone_api/internal/db"
	"ema_sound_clone_api/internal/utils/auth"
	"ema_sound_clone_api/internal/utils/crypter"
	"ema_sound_clone_api/internal/utils/random"
	"ema_sound_clone_api/pkg/model/entity"
	"ema_sound_clone_api/pkg/model/request"
	"ema_sound_clone_api/pkg/model/response"
	"ema_sound_clone_api/pkg/repository"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Admin interface {
	CreateUserAdminByDev(req request.AdminUserCreateByDevRequest) (res *response.AdminUserCreateResponse, err error)
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}

const (
	ErrAdminUserAlreadyExist = "admin user already exists"
)

func (admin) CreateUserAdminByDev(req request.AdminUserCreateByDevRequest) (res *response.AdminUserCreateResponse, err error) {
	var (
		d   = db2.GetDb()
		rp  = repository.NewAdmin()
		now = time.Now()
	)

	if adminUser, _ := rp.FindBySignIn(d, req.Email); adminUser.Email != "" {
		return nil, errors.New(ErrAdminUserAlreadyExist)
	}

	err = d.Transaction(func(tx *gorm.DB) error {
		pw := random.String(10)
		hashed, err := crypter.EncryptToHexString(pw)
		if err != nil {
			return err
		}

		issue := auth.Issue(now)

		doc := &entity.AdminUser{
			Account: &entity.Account{
				Password:   hashed,
				SignedUpAt: now,
				RefreshToken: &entity.AccountRefreshToken{
					Token:         issue.RefreshToken,
					AccessTokenID: issue.ID.String(),
					IssuedAt:      issue.IssueAt,
				},
			},
			Email: req.Email,
			Name:  req.Name,
		}

		if err := rp.Create(tx, doc); err != nil {
			return err
		}

		res = &response.AdminUserCreateResponse{
			Email:    req.Email,
			Name:     req.Name,
			Password: pw,
		}

		return nil
	})

	return res, err
}
