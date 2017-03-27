package mc

import (
	"errors"

	"github.com/qiuchengw/go-user/config"
)

func getServerList() ([]string, error) {
	servers := config.ConfigData.MemcacheServerList
	if len(servers) == 0 {
		return nil, errors.New("empty memcache server list")
	}
	return servers, nil
}
