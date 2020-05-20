package cache

import (
	_ "red-east/dao/cache/cacheimp_real"
	"red-east/dao/cache/cacheinterface"
	"red-east/dao/cache/driver"
)

func Driver(cacheType string) cacheinterface.CacheInterface {
	return driver.Driver(cacheType)
}
