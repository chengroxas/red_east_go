package cache

import (
	"fmt"
	. "red-east/utils"

	"github.com/go-redis/redis"
)

type Redis struct {
	RedisHandle *redis.Client
	CacheInterface
}

func (self *Redis) Connect() error {
	redisConfig := Config.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	self.RedisHandle = redisClient
	return nil
}

func (self *Redis) SetCache(key string, value string) error {
	result := self.RedisHandle.Set(key, value, 0)
	if result.Err() != nil {
		fmt.Println(result.Err().Error())
	}
	return nil
}
