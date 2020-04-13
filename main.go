package main

import (
	// "fmt"
	"red-east/config"

	"log"
	. "red-east/utils"

	// "reflect"
	"red-east/dao/database"
	"red-east/utils/logging"
)

func main() {
	//先获取配置
	Init()
	Logger.Info("start application....")
}

func Init() {
	var err error
	Config, err = config.InitConfig()
	if err != nil {
		log.Fatalln("init config false")
	}
	Logger = logging.InitLogger()
	if err != nil {
		log.Fatalln("init logger false")
	}

	DB, err = mysql.InitMySql()
	if err != nil {
		Logger.Error("init database false", err.Error())
	}
	defer DB.Close()
}
