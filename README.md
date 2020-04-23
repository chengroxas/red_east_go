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
|		|-redis.go
|	|-database
|		|-database.go
|	|-doc.go
|-middleware		//中间件
|  	|-middleware.go
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

3.目前只集成redis和mysql，不支持多个数据库切换以及多个缓存或者主从配置。

4.日志如果不想记录到文件中，配置logging.file_write设置为false

5.utils/common.go里包含Logger，Config，DB，Cache，Request;使用到这些必须import . "utils/common.go"