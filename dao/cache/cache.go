package cache

import (
	"fmt"
)

type CacheImp struct {
	Handle CacheInterface
}

func (self *CacheImp) InitCache() error {
	err := self.Handle.Connect()
	fmt.Printf("%+v", self.Handle)
	if err != nil {
		return err
	}
	return nil
}

func (self *CacheImp) SetCache(key string, value string) error {
	self.Handle.SetCache(key, value)
	return nil
}
