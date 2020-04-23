package cache

type CacheInterface interface {
	Connect() error
	SetCache(key string, value string) error
	GetCache(key string) string
	KeyExist(key string) bool
}
