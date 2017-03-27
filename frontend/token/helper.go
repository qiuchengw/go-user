package token

import (
	"github.com/qiuchengw/go-user/util"
)

func NewTokenId() string {
	return string(util.NewRandom())
}

func ExpirationAccess(timestamp int64) int64 {
	return timestamp + 7200
}

func ExpirationRefresh(timestamp int64) int64 {
	return timestamp + 31556952
}
