package middleware

import (
	"fmt"
	"red-east/controller"
	. "red-east/utils"

	"github.com/gin-gonic/gin"
)

//校验签名
func CheckSign() gin.HandlerFunc {

	return func(c *gin.Context) {
		// deviceType := c.MustGet("device_type")
		// sign := c.MustGet("sign")
		// t := c.MustGet("t")
		// version := c.MustGet("version")
		signConfig := Config.Sign
		path := c.Request.URL.Path
		if signConfig.Check {
			fmt.Println(path)
		}
		c.Next()
	}
}

func CheckCommonParam() gin.HandlerFunc {

	return func(c *gin.Context) {
		deviceType := c.Query("device_type")
		t := c.Query("t")
		version := c.Query("version")
		sign := c.Query("sign")
		if deviceType == "" || t == "" || version == "" || sign == "" {
			controller.Wrong(c, CODE_BAD_PARAM)
			return
		}
		c.Set("device_type", deviceType)
		c.Set("t", t)
		c.Set("sign", sign)
		c.Set("version", version)
		c.Next()
		//调完组件后要处理什么
	}
}

//校验登录
func CheckLogin() {

}

//校验用户身份 游客或者已经登录
func CheckUserMode() {

}
