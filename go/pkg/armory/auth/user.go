package auth

import (
	"crypto/sha256"
	"encoding/base64"
)

func (x *User) CheckPassword(password string) bool {
	if x != nil {
		sum := sha256.New().Sum([]byte(password + x.Salt))
		coded := base64.StdEncoding.EncodeToString(sum)
		return x.Password == coded
	}
	return false
}
