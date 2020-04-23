package cache

import (
	. "red-east/utils"
)

//根据配置返回缓存驱动
var driver CacheInterface

func Driver() CacheInterface {
	if Config.Cache.Type == "redis" {
		driver = &Redis{}
	} else {
		driver = &Memcache{}
	}
	return driver
}
