package crypter

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 10
)

func Encrypt(bytes []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(bytes, cost)
}

func Compare(hashed []byte, s []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, []byte(s))
}

func EncryptToHexString(s string) (string, error) {
	b, err := Encrypt([]byte(s))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func CompareWithHexString(hexStr string, s string) error {
	hashed, err := hex.DecodeString(hexStr)
	if err != nil {
		return err
	}

	return Compare(hashed, []byte(s))
}
