package securitykey

import (
	"errors"

	"github.com/qiuchengw/go-user/config"
)

// 从私密存储上获取安全key
func getKey() ([]byte, error) {
	key := config.ConfigData.SecurityKey
	if len(key) != 128 {
		return nil, errors.New("incorrect security key")
	}
	return key, nil
}
