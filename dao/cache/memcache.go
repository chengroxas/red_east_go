package cache

import (
	"red-east/minterface"
)

//todo 实现memcache
type Memcache struct {
	minterface.CacheInterface
}

func (self *Memcache) Connect() error {
	return nil
}
