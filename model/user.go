package model

import (
	. "red-east/utils"
)

type UserModel struct {
	Id       int
	UserId   int64
	Username string
}

/**
 * 需要设置表名，不然会根据结构体的名称默认表名user_model
 */
func (UserModel) TableName() string {
	return Config.MySql.Prefix + "user"
}

func (self *UserModel) GetOneInfo(query interface{}, args ...interface{}) (notFound bool) {
	notFound = DB.Where(query, args).First(&self).RecordNotFound()
	return
}

func (self *UserModel) GetUserList(page, pageSize int, query interface{}, args ...interface{}) ([]*UserModel, int) {
	var userList []*UserModel
	var count int
	DB.Where(query, args).Offset(page).Limit(pageSize).Find(&userList)
	DB.Model(UserModel{}).Count(&count)
	return userList, count
}
