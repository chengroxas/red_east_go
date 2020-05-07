package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"red-east/config"
	"red-east/dao/cache"
	"red-east/dao/database"
	"red-east/imp"
	"red-east/logging"
	"red-east/router"
	. "red-east/utils"
)

func main() {
	var err error
	//初始化配置
	Config, err = config.InitConfig()
	if err != nil {
		log.Fatalln("init config fail", err.Error())
	}
	//初始化日志
	Logger, err = logging.InitLogger()
	if err != nil {
		log.Fatalln("init logger fail", err.Error())
	}
	//初始化数据库，使用mysql
	DB, err = database.InitMySql()
	if err != nil {
		Logger.Error("init database fail", err.Error())
	}
	defer DB.Close()
	//初始化缓存
	cacheDriver := cache.Driver()
	Cache = imp.CacheImp{
		Handle: cacheDriver,
	}
	cacheErr := Cache.InitCache(Config, Logger)
	if cacheErr != nil {
		Logger.Error("init cache fail:", cacheErr.Error())
	}
	defer Cache.Close()
	// 初始化gin
	// gin输出到文件或者终端
	gin.DefaultWriter = io.MultiWriter(logging.Writers...)
	//去掉颜色
	gin.DisableConsoleColor()
	r := gin.New()
	//中间件要在路由注册之前
	r.Use(gin.Logger())

	//注册路由
	router.RegisterRouter(r)
	//使用中间件，记录请求
	Logger.Info("start application....")
	r.Run()
}
