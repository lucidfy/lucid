package hash

import "golang.org/x/crypto/bcrypt"

func Make(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	return string(hashed), err
}

func Check(plainText string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText)) == nil
}
