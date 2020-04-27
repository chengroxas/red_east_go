package cache

import (
	"red-east/minterface"
	"red-east/utils"
)

//根据配置返回缓存驱动
var driver minterface.CacheInterface

func Driver() minterface.CacheInterface {
	if common.Config.Cache.Type == "redis" {
		driver = &Redis{}
	} else {
		driver = &Memcache{}
	}
	return driver
}
