package cache

type Memcache struct {
	CacheInterface
}

func (self *Memcache) Connect() error {
	return nil
}
