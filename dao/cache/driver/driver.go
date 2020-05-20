package driver

import (
	"red-east/dao/cache/cacheinterface"
)

//根据配置返回缓存驱动
//var driver minterface.CacheInterface

var driver = make(map[string]cacheinterface.CacheInterface)

func Register(cacheType string, cache cacheinterface.CacheInterface) {
	driver[cacheType] = cache
}

func Driver(cacheType string) cacheinterface.CacheInterface {
	if value, exist := driver[cacheType]; exist {
		return value
	} else {
		panic("没有某个类的驱动")
	}
	//if common.Config.Cache.Type == "redis" {
	//	driver = &Redis{}
	//} else {
	//	driver = &Memcache{}
	//}
	//return driver
}
