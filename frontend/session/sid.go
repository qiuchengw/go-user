package session

import (
	"github.com/qiuchengw/go-user/util"
)

// ^[A-Za-z0-9_-]+$
func NewSessionId() (sid string, err error) {
	return util.NewID()
}

// ^temp\.[A-Za-z0-9_-]+$
func NewGuestSessionId() (sid string, err error) {
	sidx, err := NewSessionId()
	if err != nil {
		return
	}
	sid = "temp." + string(sidx)
	return
}
