package hash

import (
	"github.com/lucidfy/lucid/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Make(plainText string) (string, *errors.AppError) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	return string(hashed), errors.InternalServerError("bcrypt.GenerateFromPassword() error", err)
}

func Check(plainText string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText)) == nil
}
