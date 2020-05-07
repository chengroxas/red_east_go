package common

import (
	"fmt"
	"red-east/imp"
	"strconv"

	"github.com/gin-gonic/gin"

	// "fmt"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"red-east/config"

	"red-east/logging"
	"red-east/utils/external"

	"github.com/jinzhu/gorm"
)

var (
	Logger  logging.NLogger
	Config  config.Config
	DB      *gorm.DB
	Cache   imp.CacheImp
	Request external.Request
)

func GetPageParam(c *gin.Context) (page int, pageSize int) {
	var err error
	page, err = strconv.Atoi(c.Query("page"))
	pageSize, err = strconv.Atoi(c.Query("page_size"))
	page = (page - 1) * pageSize
	if err != nil {
		return 0, 10
	}
	return page, pageSize
}

func Md5ToString(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)
	return md5str
}

func Sha256ToString(str string) string {
	data := []byte(str)
	hash := sha256.Sum256(data)
	sha256str := fmt.Sprintf("%x", hash)
	return sha256str
}

func Base64ToString(str string) string {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}

// func GetStructFieldValue(data interface{}, field string) {
// 	rs := reflect.ValueOf(data)
// 	fmt.Printf("%+v", rs)
// }
