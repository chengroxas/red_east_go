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

type ResUserList struct {
	List []User `json:"list"`
	Cnt  int    `json:"cnt"`
}

type User struct {
	Id       int    `json:"id"`
	UserId   int64  `json:"user_id"`
	UserName string `json:"username"`
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
	//list := new([]UserList) //使用new，当数据为空的时候，list为nil
	list := make([]User, 0) //使用make，当数据为空的时候，list为[]
	page, pageSize := GetPageParam(c)
	usersData, count := userModel.GetUserList(page, pageSize, "country_code = ?", "86")
	for _, userData := range usersData {
		user := &User{
			Id:       userData.Id,
			UserId:   userData.UserId,
			UserName: userData.Username + "qly",
		}
		list = append(list, *user)
	}
	res := ResUserList{
		List: list,
		Cnt:  count,
	}
	self.Success(c, res)
}
