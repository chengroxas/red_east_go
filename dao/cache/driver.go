package cache

import (
	. "red-east/utils"
)

var driver CacheInterface

func Driver() CacheInterface {
	if Config.Cache.Type == "redis" {
		driver = &Redis{}
	} else {
		driver = &Memcache{}
	}
	return driver
}
