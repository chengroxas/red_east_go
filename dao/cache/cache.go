package cache

import (
	"fmt"
	. "red-east/utils"
)

type Cache struct {
	Handle interface{}
}

func InitCache() {
	cacheConfig := Config.Cache
	fmt.Println(cacheConfig)
}
