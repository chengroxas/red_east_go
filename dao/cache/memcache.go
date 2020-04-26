package cache

//todo 实现memcache
type Memcache struct {
	CacheInterface
}

func (self *Memcache) Connect() error {
	return nil
}
