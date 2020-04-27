package imp

// "fmt"
import (
	"red-east/config"
	"red-east/logging"
	"red-east/minterface"
	"time"
)

//缓存实现类，这里处理想要做的事情，例如加日志或者加key前缀
type CacheImp struct {
	Handle minterface.CacheInterface
	Config config.Config
	Logger logging.NLogger
}

func (self *CacheImp) InitCache(config config.Config, logger logging.NLogger) error {
	err := self.Handle.Connect()
	self.Config = config
	self.Logger = logger
	if err != nil {
		return err
	}
	return nil
}

func (self *CacheImp) SetCache(key string, value string, expire time.Duration) error {
	prefix := self.Config.Cache.Prefix
	key = prefix + key
	err := self.Handle.SetCache(key, value, expire*time.Second)
	if err != nil {
		self.Logger.Error("cache set key: ", key, " fail:", err.Error())
		return err
	}
	return nil
}

func (self *CacheImp) GetCache(key string) (string, error) {
	prefix := self.Config.Cache.Prefix
	key = prefix + key
	value, err := self.Handle.GetCache(key)
	if err != nil {
		self.Logger.Error("cache get key ", key, " fail:", err.Error())
		return "", nil
	}
	return value, nil
}

func (self *CacheImp) KeyExist(key string) (bool, error) {
	prefix := self.Config.Cache.Prefix
	key = prefix + key
	exist, err := self.Handle.KeyExist(key)
	if err != nil {
		self.Logger.Error("cache key ", key, " exist fail:", err.Error())
		return false, err
	}
	return exist, nil
}

func (self *CacheImp) Close() error {
	err := self.Handle.Close()
	if err != nil {
		self.Logger.Error("cache close fail:", err.Error())
	}
	return err
}
