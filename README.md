### golang demo框架

gin + gorm + redis/memcache + rotatelogs(日志处理) + go-nsq(消息队列)

目录结构

```go
|-config
|	|-config.go
|-controller
|	|-controller.go		//controller基类
|	|-user
|		|-tourist.go
|		|-user.go
|-dao						//数据层
|---cache
|    |--cacheimp             //这个cache实现业务层的逻辑
|      |--cache_imp.go
|    |--cacheimp_real        //只需实现第三方缓存相关功能即可
|      |--memcache.go        
|      |--redis.go
|    |--cacheinterface      //cache接口，cacheimp和cacheimp_real必须实现
|      |--cache_interface.go
|    |--driver              //驱动，注册cache
|      |--driver.go 
|    |--cache.go //给外部提供Driver方法调用driver.Driver方法，避免包重复
|---database
|	 |-database.go
|  |-doc.go
|-middleware		//中间件
|  	|-middleware.go
|-queue        //队列，用的nsq
|   |-mconsumer
|     |-sms_consumer.go //消费者
|   |-mproducer       //生产者放到这里
|     |-sms_producer.go  //示例
|     |-producer.go   //nsq producer
|   |- queue.go       //开启消费
|-router					//路由
|	|-router.go
|-service					//服务
|	|-service.go
|-utils						//工具类
|	|-external
|		|-request.go
|	|-logging				//日志
|		|-logging.go
|	|-common.go			//公共类
|	|-errors.go				//错误码定义
|-config.yaml			//配置
|-main.go
|-README.md
```

### 注意事项

1.请开启gomodule模式，在根目录下运行go mode init，在根目录下运行go run main.go运行程序，调试时可以使用bee run，自动重新编译运行

2.开头配置一定要先初始化；配置文件在根目录下的config.yaml，如果添加新的配置，则在config/config.go中定义好结构体。

3.缓存可以使用redis和memcache，可以切换但是不能混合使用。如要添加新的缓存类型或者新的方法，查看dao/cache/README.md。

4.orm使用gorm。

5.日志在logging目录下，使用file-rotatelogs进行文件切割。如果日志记录不想写到文件中，修改配置文件logging.file_write设置为false。

6.utils/common.go里定义了全局变量Logger，Config，DB，Cache，Request;在main.go中进行初始化，
使用时，import . "utils/common.go"

7.队列使用go-nsq，请提起安装好golang的nsq。

8.