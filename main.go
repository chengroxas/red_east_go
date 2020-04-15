package main

import (
	"github.com/gin-gonic/gin"

	// "fmt"
	"red-east/config"

	"log"
	. "red-east/utils"

	// "reflect"
	"red-east/dao/database"
	"red-east/utils/logging"
	// "github.com/gin-gonic/gin"
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
		log.Fatalln("init config fail")
	}
	Logger = logging.InitLogger()
	if err != nil {
		log.Fatalln("init logger fail")
	}

	DB, err = mysql.InitMySql()
	if err != nil {
		Logger.Error("init database fail", err.Error())
	}
	defer DB.Close()
	r := gin.New()
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, "hwllo worj")
	})
	r.Run()
}
