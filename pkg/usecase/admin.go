package usecase

import (
	"ema_sound_clone_api/config"
	"ema_sound_clone_api/internal/db"
	"ema_sound_clone_api/internal/utils/auth"
	"ema_sound_clone_api/internal/utils/crypter"
	"ema_sound_clone_api/internal/utils/random"
	"ema_sound_clone_api/pkg/model/entity"
	"ema_sound_clone_api/pkg/model/request"
	"ema_sound_clone_api/pkg/model/response"
	"ema_sound_clone_api/pkg/repository"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Admin interface {
	CreateUserAdminByDev(req request.AdminUserCreateByDevRequest) (res *responsemodel.AdminUserCreateResponse, err error)
	SignIn(req request.AdminUserLoginRequest) (res *responsemodel.AdminLoginResponse, err error)
	RefreshToken(req request.AdminUserRefreshToken) (res *responsemodel.AdminLoginResponse, err error)
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}

const (
	ErrAdminUserAlreadyExist = "admin user already exists"
	ErrAdminUserNotFound     = "admin user not found"
	ErrAccountAdminUserIsNil = "admin user account is nill"
	ErrRefreshTokenInvalid   = "invalid refresh token"
	ErrInvalidPassword       = "invalid email or password"
)

func (admin) CreateUserAdminByDev(req request.AdminUserCreateByDevRequest) (res *responsemodel.AdminUserCreateResponse, err error) {
	var (
		d   = db.GetDb()
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

		res = &responsemodel.AdminUserCreateResponse{
			Email:    req.Email,
			Name:     req.Name,
			Password: pw,
		}

		return nil
	})

	return res, err
}

func (admin) SignIn(req request.AdminUserLoginRequest) (res *responsemodel.AdminLoginResponse, err error) {
	var (
		d   = db.GetDb()
		rp  = repository.NewAdmin()
		now = time.Now()
		env = config.GetEnv()
	)

	adminUser, err := rp.FindBySignIn(d, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(ErrAdminUserNotFound)
		}
		return nil, err
	}

	err = d.Transaction(func(tx *gorm.DB) error {
		var err error

		if err = crypter.CompareWithHexString(adminUser.Account.Password, req.Password); err != nil {
			return errors.New(ErrInvalidPassword)
		}

		token := auth.Issue(now)

		// assign new token
		if adminUser.Account == nil || adminUser.Account.RefreshToken == nil {
			return errors.New(ErrAccountAdminUserIsNil)
		}
		adminUser.Account.RefreshToken.Token = token.RefreshToken
		adminUser.Account.RefreshToken.IssuedAt = token.IssueAt
		adminUser.Account.RefreshToken.AccessTokenID = token.ID.String()

		if err = rp.UpdateRefreshToken(d, adminUser); err != nil {
			return err
		}

		claims := auth.NewClaims(token.ID, strconv.Itoa(int(adminUser.Account.ID)), adminUser.Email, now)

		accessToken, err := auth.Sign(env.AdminJWTKey, &claims)
		if err != nil {
			return err
		}

		res = &responsemodel.AdminLoginResponse{
			AccessToken:  accessToken,
			RefreshToken: token.RefreshToken,
		}

		return nil
	})

	return res, err
}

func (admin) RefreshToken(req request.AdminUserRefreshToken) (res *responsemodel.AdminLoginResponse, err error) {
	var (
		d   = db.GetDb()
		rp  = repository.NewAdmin()
		now = time.Now()
		env = config.GetEnv()
	)

	adminUser, err := rp.FindByRefreshToken(d, req.RefreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(ErrAdminUserNotFound)
		}
		return nil, err
	}

	if adminUser.Account.RefreshToken.Expired(now) {
		return nil, errors.New(ErrRefreshTokenInvalid)
	}

	err = d.Transaction(func(tx *gorm.DB) error {
		var err error

		token := auth.Issue(now)

		// assign new token
		if adminUser.Account == nil || adminUser.Account.RefreshToken == nil {
			return errors.New(ErrAccountAdminUserIsNil)
		}
		adminUser.Account.RefreshToken.Token = token.RefreshToken
		adminUser.Account.RefreshToken.IssuedAt = token.IssueAt
		adminUser.Account.RefreshToken.AccessTokenID = token.ID.String()

		if err = rp.UpdateRefreshToken(d, adminUser); err != nil {
			return err
		}

		claims := auth.NewClaims(token.ID, strconv.Itoa(int(adminUser.Account.ID)), adminUser.Email, now)

		accessToken, err := auth.Sign(env.AdminJWTKey, &claims)
		if err != nil {
			return err
		}

		res = &responsemodel.AdminLoginResponse{
			AccessToken:  accessToken,
			RefreshToken: token.RefreshToken,
		}

		return nil
	})

	return res, err
}
