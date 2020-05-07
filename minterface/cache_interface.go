package minterface

import (
	"time"
)

//单纯的接口，不是实现任何东西
type CacheInterface interface {
	Connect() error
	SetCache(key string, value string, expire time.Duration) error
	GetCache(key string) (string, error)
	KeyExist(key string) (bool, error)
	Close() error
}
