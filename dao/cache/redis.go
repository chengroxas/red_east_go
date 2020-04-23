package cache

import (
	. "red-east/utils"
	"time"

	"github.com/go-redis/redis"
)

//这里只需要实现redis，无需考虑别的操作，但是这里必须实现CacheInterface里的方法
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

func (self *Redis) SetCache(key string, value string, expire time.Duration) error {
	result := self.Handle.Set(key, value, expire)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (self *Redis) GetCache(key string) (string, error) {
	result, err := self.Handle.Get(key).Result()
	if err == redis.Nil {
		return result, nil
	} else if err != nil {
		return "", err
	}
	return result, nil
}

func (self *Redis) KeyExist(key string) (bool, error) {
	result := self.Handle.Exists(key)
	if result.Err() != nil {
		return false, result.Err()
	}
	if result.Val() == 0 {
		return false, nil
	}
	return true, nil
}

func (self *Redis) Close() error {
	return self.Handle.Close()
}
