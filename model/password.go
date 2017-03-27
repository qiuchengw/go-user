package model

import (
	"crypto/hmac"
	"crypto/sha1"

	"github.com/qiuchengw/go-user/util"
)

var PasswordSalt = NewSalt() // 无主地 salt, 用于安全授权

func NewSalt() []byte {
	rd := util.NewRandom()
	return []byte(rd)
}

func NewPasswordTag() string {
	return string(util.NewRandom())
}

func EncryptPassword(password, salt []byte) []byte {
	Hash := hmac.New(sha1.New, salt)
	Hash.Write(password)
	return Hash.Sum(nil)
}
