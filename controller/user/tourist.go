package user

import (
	"github.com/gin-gonic/gin"
)

type Tourist struct{}

func (this *Tourist) Login(c *gin.Context) {
	c.JSON(200, "user tourist login")
}
