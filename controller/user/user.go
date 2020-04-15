package user

import (
	"github.com/gin-gonic/gin"
)

type User struct{}

func (this *User) LoginByPass(c *gin.Context) {
	c.JSON(200, "user user login")
}

func (this *User) LoginBySms(c *gin.Context) {
	c.JSON(200, "user user register")
}
