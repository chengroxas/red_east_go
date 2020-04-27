package model

import (
	. "red-east/utils"
)

type UserModel struct {
	Id     int
	UserId int64
}

func (UserModel) TableName() string {
	return "user"
}

func (self *UserModel) GetOneInfo(query interface{}, args ...interface{}) bool {
	notFound := DB.Where(query, args).First(&self).RecordNotFound()
	return notFound
}

func (self *UserModel) GetUserList(page, pageSize int, query interface{}, args ...interface{}) ([]*UserModel, int) {
	var userList []*UserModel
	var count int
	DB.Where(query, args).Limit(pageSize).Offset(page).Find(&userList)
	return userList, count
}
