package middleware

import (
	"fmt"
	"red-east/controller"
	. "red-east/utils"
	"reflect"
	"strconv"
	"strings"
	"time"

	// "time"

	"github.com/gin-gonic/gin"
)

//校验签名
func CheckSign() gin.HandlerFunc {

	return func(c *gin.Context) {
		deviceType := c.GetString("device_type")
		sign := c.GetString("sign")
		t := c.GetString("t")
		signConfig := Config.Sign
		path := c.Request.URL.Path
		if signConfig.Check {
			//通过反射获取结构体AppKey里的值 deviceType IOS|Web|Andorid
			refelctAppKey := reflect.ValueOf(signConfig.AppKey)
			appKeyValue := refelctAppKey.FieldByName(deviceType)
			appKey := appKeyValue.Interface().(string)
			//获取时间戳
			timeInt64, _ := strconv.ParseInt(t, 10, 64)
			timeStamp := timeInt64 / 1000
			nowTimeStamp := time.Now().Unix()
			if nowTimeStamp-timeStamp > 300 {
				controller.Wrong(c, CODE_BAD_AUTH)
				return
			}
			//这样拼接效率会更快 path+deviceType+appKey+t
			var strBuffer strings.Builder
			strBuffer.WriteString(path)
			strBuffer.WriteString(deviceType)
			strBuffer.WriteString(appKey)
			strBuffer.WriteString(t)
			s := strBuffer.String()
			//加密签名
			signReal := Md5ToString(s)

			signParam := strings.ToLower(sign)
			signReal = strings.ToLower(signReal)
			if signParam != signReal {
				Logger.Error("客户端的签名:", signParam, "!=真实的签名", signReal)
				controller.Wrong(c, CODE_BAD_AUTH)
				return
			}
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
		//这里必须要判断deviceType的值，因为CheckSign通过反射字段获得值，字段必须存在
		if deviceType != "IOS" && deviceType != "Web" && deviceType != "Android" {
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
