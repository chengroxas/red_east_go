### golang demo框架

gin + mysql + redis + rotatelogs(日志处理)

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
|	|-cache
|       |-cache_driver.go //缓存的驱动
|       |-memcache.go
|		|-redis.go     //缓存的具体实现
|	|-database
|		|-database.go
|	|-doc.go
|-imp
|   |-cache_imp.go   //缓存的实现
|-minterface
|   |-cache_interface.go //缓存的接口
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

2.配置必须第一个初始化；如果添加配置，必须修改config.go，定义相关结构体使配置生效。

3.缓存可以使用redis和memcache，可以切换。如需其它缓存，则需要在cache中编写类实现cache_interface接口，
config.yaml里加入对应配置，config/config.go定义配置的结构体，cache_driver.go里返回相应缓存结构体。
不支持混合使用。

4.使用的是gorm处理数据模型。

5.日志如果不想记录到文件中，配置logging.file_write设置为false

6.utils/common.go里包含Logger，Config，DB，Cache，Request;使用到这些必须import . "utils/common.go"