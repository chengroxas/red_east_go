package cache_type

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"red-east/dao/cache/driver"
	. "red-east/utils"
	"time"
)

type Memcache struct {
	Handle *memcache.Client
	//minterface.CacheInterface
}

func init() {
	cache := &Memcache{}
	driver.Register("memcache", cache)
}

func (self *Memcache) Connect() error {
	dns := fmt.Sprintf("%s:%s", Config.Memcache.Host, Config.Memcache.Port)
	memcacheClient := memcache.New(dns)
	err := memcacheClient.Ping()
	if err != nil {
		return err
	}
	self.Handle = memcacheClient
	return nil
}

func (self *Memcache) SetCache(key string, value string, expire time.Duration) error {
	err := self.Handle.Set(&memcache.Item{Key: key, Value: []byte(value), Flags: 0, Expiration: int32(expire)})
	return err
}

func (self *Memcache) GetCache(key string) (string, error) {
	item, err := self.Handle.Get(key)
	if err != nil && err != memcache.ErrCacheMiss {
		return "", err
	}
	return string(item.Value), nil
}

func (self *Memcache) KeyExist(key string) (bool, error) {
	_, err := self.Handle.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (self *Memcache) Close() error {
	return nil
}
