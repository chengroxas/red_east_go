package cache

import (
	. "red-east/utils"

	"github.com/go-redis/redis"
)

type Redis struct {
	Handle *redis.Client
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
	self.Handle = redisClient
	return nil
}

func (self *Redis) SetCache(key string, value string) error {
	self.Handle.Set(key, value, 0)
	return nil
}
