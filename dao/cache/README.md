缓存的实现

cache配置文件config.yaml，添加配置要在config/config.go里定义对应的结构体

```yaml
cache:
  type: redis
redis:
  host: 10.11.1.90
  port: 6379
  db: 1
  password: 
  prefix: RED_EAST_
memcache:
	##memcache的配置，因为memcache少用就不做了
```

缓存cache的实现

```go
|---cache
|  |--cacheimp             //这个cache实现业务层的逻辑
|    |--cache_imp.go
|  |--cacheimp_real        //只需实现第三方缓存相关功能即可
|    |--memcache.go        
|    |--redis.go
|  |--cacheinterface      //cache接口，cacheimp和cacheimp_real必须实现
|    |--cache_interface.go
|  |--driver              //驱动，注册cache
|    |--driver.go
|  |--cache.go            //给外部提供Driver方法调用driver.Driver方法，避免包重复
```

1.如果要添加新方法，cacheimp和cacheinterface，cacheimp_real里都必须实现新的方法
2.如果添加新的缓存类型，只需要在cacheimp_real添加并且实现cacheimp里的方法，
并且init方法里必须调用driver.Register注册到驱动中
