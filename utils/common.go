package common

import (
	"fmt"
	// "fmt"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"red-east/config"

	"red-east/utils/external"
	"red-east/utils/logging"

	"github.com/jinzhu/gorm"
	"red-east/minterface"
)

var (
	Logger   logging.NLogger
	Config   config.Config
	DB       *gorm.DB
	Cache    minterface.CacheInterface
	Request  external.Request
	Page     int
	PageSize int
)

//func GetPageParam() (page, pageSize int) {
//
//}

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
