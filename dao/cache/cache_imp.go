package cache

// "fmt"
import (
	"red-east/minterface"
	. "red-east/utils"
	"time"
)

//缓存实现类，这里处理想要做的事情，例如加日志或者加key前缀
type CacheImp struct {
	Handle minterface.CacheInterface
}

func (self *CacheImp) InitCache() error {
	err := self.Handle.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (self *CacheImp) SetCache(key string, value string, expire time.Duration) error {
	prefix := Config.Cache.Prefix
	key = prefix + key
	err := self.Handle.SetCache(key, value, expire*time.Second)
	if err != nil {
		Logger.Error("cache set key: ", key, " fail:", err.Error())
		return err
	}
	return nil
}

func (self *CacheImp) GetCache(key string) (string, error) {
	prefix := Config.Cache.Prefix
	key = prefix + key
	value, err := self.Handle.GetCache(key)
	if err != nil {
		Logger.Error("cache get key ", key, " fail:", err.Error())
		return "", nil
	}
	return value, nil
}

func (self *CacheImp) KeyExist(key string) (bool, error) {
	prefix := Config.Cache.Prefix
	key = prefix + key
	exist, err := self.Handle.KeyExist(key)
	if err != nil {
		Logger.Error("cache key ", key, " exist fail:", err.Error())
		return false, err
	}
	return exist, nil
}

func (self *CacheImp) Close() error {
	err := self.Handle.Close()
	if err != nil {
		Logger.Error("cache close fail:", err.Error())
	}
	return err
}
