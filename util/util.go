package util

import (
	"time"

	"github.com/anarcher/shortuuid"
)

var _used int64 = 0

// NewID generate a random / unique string
func NewID() (id string, err error) {
	return NewRandom(), nil
}

// NewRandom gen random string
func NewRandom() string {
	alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxy"
	uid := shortuuid.NewWithAlphabet(alphabet) // uBFWRLr5dXbeWfiasZi
	return uid.String()
}

func Int64ID() int64 {
	t := time.Now()
	_used++
	return t.UnixNano() + _used
}
