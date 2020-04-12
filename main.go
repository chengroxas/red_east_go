package main

import (
	"red-east/config"
	"red-east/utils/logging"
)

func main() {
	//获取配置
	err := config.InitConfig()
	if err != nil {
	}
	//日志，可能需要输入到某个文件中
	logger := logging.InitLogger()
	logger.Info("hello world")
	//db连接
}
