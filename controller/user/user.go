package user

import (
	"github.com/gin-gonic/gin"
	"red-east/controller"
	"red-east/model"
	. "red-east/utils"
)

type UserController struct {
	controller.Base
}

type RevLoginDataType struct {
	Mobile string `form:"mobile" binding:"required"`
}

func (self *UserController) LoginByPass(c *gin.Context) {
	var args RevLoginDataType
	if err := c.ShouldBind(&args); err != nil {
		self.Wrong(c, CODE_BAD_PARAM)
		return
	}
	userModel := model.UserModel{}
	notFound := userModel.GetOneInfo("mobile = ?", args.Mobile)
	if notFound {
		self.Wrong(c, CODE_NOT_EXIST)
		return
	}
	self.Success(c, userModel)
}

func (self *UserController) LoginBySms(c *gin.Context) {
	c.JSON(200, "user user register")
}

func (self *UserController) GetUserList(c *gin.Context) {
	userModel := model.UserModel{}
	userList, count := userModel.GetUserList(1, 10, "country_code = ?", "86")
	Logger.Info(userList)
	Logger.Info("总数:", count)
}
