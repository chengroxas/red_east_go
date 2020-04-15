package main

import (
	"io"
	"log"

	"red-east/config"
	"red-east/dao/database"
	"red-east/middleware"
	"red-east/router"
	. "red-east/utils"
	"red-east/utils/logging"

	"github.com/gin-gonic/gin"
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
	DB, err = mysql.InitMySql()
	if err != nil {
		Logger.Error("init database fail", err.Error())
	}
	defer DB.Close()
	// 初始化gin
	// gin输出到文件或者终端
	gin.DefaultWriter = io.MultiWriter(logging.Writers...)
	//去掉颜色
	gin.DisableConsoleColor()
	r := gin.New()
	//中间件要在路由注册之前
	r.Use(gin.Logger(), middleware.CheckCommonParam(), middleware.CheckSign())

	//注册路由
	router.RegisterRouter(r)
	//使用中间件，记录请求
	Logger.Info("start application....")

	r.Run()
}
