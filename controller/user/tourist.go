package user

import (
	. "red-east/controller"

	"github.com/gin-gonic/gin"
)

type Tourist struct {
	Base
}

func (this *Tourist) Login(c *gin.Context) {
	Success(c, nil)
}
