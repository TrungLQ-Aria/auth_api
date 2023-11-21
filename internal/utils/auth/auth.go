package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
}

type Token struct {
	ID           uuid.UUID
	RefreshToken string
	IssueAt      time.Time
}

const (
	Expiration = time.Hour * 24
)

func NewClaims(id uuid.UUID, accountID string, signInID string, now time.Time) Claims {
	return Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(Expiration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "auth-admin",
			Subject:   accountID,
			ID:        id.String(),
			Audience:  []string{signInID},
		},
	}
}

func Issue(now time.Time) Token {
	return Token{
		ID:           uuid.New(),
		RefreshToken: uuid.New().String(),
		IssueAt:      now,
	}
}

func Sign(key string, c *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(key))
}

func UnSign(key string, sign string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(sign, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("access token parse failure")
}
