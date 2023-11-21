package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func TestNewClaims(t *testing.T) {
	type args struct {
		id        uuid.UUID
		accountID string
		signInID  string
		now       time.Time
	}

	var (
		id        = uuid.New()
		accountID = "999"
		signInID  = "trungluong@gmail.com"
		now       = time.Now()
	)

	tests := []struct {
		name string
		args args
		want Claims
	}{
		{
			name: "standard",
			args: args{
				id:        id,
				accountID: accountID,
				signInID:  signInID,
				now:       now,
			},
			want: Claims{
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "auth-admin",
					Subject:   accountID,
					Audience:  []string{signInID},
					ExpiresAt: jwt.NewNumericDate(now.Add(Expiration)),
					NotBefore: jwt.NewNumericDate(now),
					IssuedAt:  jwt.NewNumericDate(now),
					ID:        id.String(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if rs := NewClaims(tt.args.id, tt.args.accountID, tt.args.signInID, tt.args.now); !reflect.DeepEqual(rs, tt.want) {
				t.Errorf("issuer.issue = %v, want %v", rs, tt.want)
			}
		})
	}
}
