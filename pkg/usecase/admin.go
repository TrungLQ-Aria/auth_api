package usecase

import (
	"errors"
	"github.com/trungaria/auth_api.git/config"
	"github.com/trungaria/auth_api.git/internal/db"
	"github.com/trungaria/auth_api.git/internal/utils/auth"
	"github.com/trungaria/auth_api.git/internal/utils/crypter"
	"github.com/trungaria/auth_api.git/internal/utils/random"
	"github.com/trungaria/auth_api.git/pkg/handler/openapi"
	"github.com/trungaria/auth_api.git/pkg/model/entity"
	"github.com/trungaria/auth_api.git/pkg/repository"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Admin interface {
	CreateUserAdminByDev(req openapi.AdminUserCreateRequest) (res *openapi.AdminUserCreateResponse, err error)
	SignIn(req openapi.AdminUserSignInRequest) (res *openapi.AdminUserSignInResponse, err error)
	RefreshToken(req openapi.PostV1AdminUserAccessTokenParams) (res *openapi.AdminUserSignInResponse, err error)
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

func (admin) CreateUserAdminByDev(req openapi.AdminUserCreateRequest) (res *openapi.AdminUserCreateResponse, err error) {
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

		res = &openapi.AdminUserCreateResponse{
			Email:    req.Email,
			Name:     req.Name,
			Password: pw,
		}

		return nil
	})

	return res, err
}

func (admin) SignIn(req openapi.AdminUserSignInRequest) (res *openapi.AdminUserSignInResponse, err error) {
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

		res = &openapi.AdminUserSignInResponse{
			AccessToken:  accessToken,
			RefreshToken: token.RefreshToken,
		}

		return nil
	})

	return res, err
}

func (admin) RefreshToken(req openapi.PostV1AdminUserAccessTokenParams) (res *openapi.AdminUserSignInResponse, err error) {
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

		res = &openapi.AdminUserSignInResponse{
			AccessToken:  accessToken,
			RefreshToken: token.RefreshToken,
		}

		return nil
	})

	return res, err
}
