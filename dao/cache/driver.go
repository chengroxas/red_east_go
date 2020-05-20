package cache

import (
	_ "red-east/dao/cache/cache_type"
	"red-east/dao/cache/driver"
	"red-east/minterface"
)

func Driver(cacheType string) minterface.CacheInterface {
	return driver.Driver(cacheType)
}
