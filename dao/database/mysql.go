package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	. "red-east/utils"
)

func InitMySql() (*gorm.DB, error) {
	//获取数据库配置
	mysqlConfig := Config.MySql
	//获取main数据库数据
	mainConfig := mysqlConfig.Main

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return mysqlConfig.Prefix + defaultTableName
	//}
	//拼接dns
	dns := "%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%ds"
	dns = fmt.Sprintf(dns, mainConfig.UserName, mainConfig.Password,
		mainConfig.Dns, mainConfig.Port, mainConfig.DataBase, mysqlConfig.Charset, mysqlConfig.ConnectTimeOut)

	Db, err := gorm.Open("mysql", dns)
	if mysqlConfig.Debug == "true" {
		Db.LogMode(true)
		if mysqlConfig.FileWrite == "true" {
			logger := Logger.GetWriter()
			if len(logger) == 2 {
				newLog := log.New(logger[1], "\r\n[DB SQL] ", 0)
				Db.SetLogger(newLog)
			}
		}
	}
	Db.SingularTable(true)
	return Db, err

}
