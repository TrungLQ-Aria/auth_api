package crypter

import (
	"testing"
)

func TestHashedPassword(t *testing.T) {
	type args struct {
		password string
	}

	tests := []struct {
		name      string
		args      args
		mustError bool
	}{
		{
			name: "success case",
			args: args{
				password: "test-password",
			},
			mustError: false,
		},
		{
			name: "failed case",
			args: args{
				password: "test-password",
			},
			mustError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := EncryptToHexString(tt.args.password)

			if err != nil {
				t.Errorf("encrypt password failed %s", tt.args.password)
				t.FailNow()
			}

			err = CompareWithHexString(hashed, tt.args.password)
			if err != nil && !tt.mustError {
				t.Errorf("hashed = %s, password = %s", hashed, tt.args.password)
			}
		})
	}
}
