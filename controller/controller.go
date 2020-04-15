package controller

import (
	. "red-east/utils"

	"github.com/gin-gonic/gin"
)

type Base struct{}

func (self *Base) Wrong(c *gin.Context, code int) {
	wrong(c, code, "")
}

func (self *Base) WrongMsg(c *gin.Context, code int, msg string) {
	wrong(c, code, msg)
}

func Wrong(c *gin.Context, code int) {
	msg := GetErrorCodeMsg(code)
	wrong(c, code, msg)
}

func WrongMsg(c *gin.Context, code int, msg string) {
	wrong(c, code, msg)
}

func wrong(c *gin.Context, code int, msg string) {
	data := map[string]string{}
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  GetErrorCodeMsg(code),
	})
	c.Abort()
}

func (self *Base) Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
	c.Abort()
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
	c.Abort()
}
