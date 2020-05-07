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
|--cache
	--driver.go						cache类的驱动
	--memcache.go			memcache缓存的实现
	--redis.go							redis缓存的实现
|--imp
    --cahce_imp.go          cache类的实现(实际调用的是各种类型的具体实现)
|--minterface
    --cache_interface.go   cache的接口
```



cache_interface.go里定义的是需要的接口。

cache_imp.go则是实现cache_interface.go里的接口，在该实现类里做额外的处理例如加日志、给key值加前缀以及对时间的处理，不需要关心redis以及memcache的具体实现。

redis.go、memcache.go则是缓存的具体实现，连接、设置、查询等，不需要做额外的处理，不需要关心别的业务逻辑。

cache_interface.go里每加一种方法，cache_imp.go以及redis、memcache或者别的缓存实现里都要实现该方法。

