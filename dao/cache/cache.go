package cache

type CacheImp struct {
	Handle CacheInterface
}

func (self *CacheImp) InitCache() error {
	err := self.Handle.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (self *CacheImp) SetCache(key string, value string) error {
	self.Handle.SetCache(key, value)
	return nil
}
