package common

import (
	"fmt"
	// "fmt"
	"crypto/md5"
	"crypto/sha256"
	"red-east/config"

	// "red-east/dao/database"
	"red-east/utils/logging"

	"github.com/jinzhu/gorm"
)

var (
	Logger logging.NLogger
	Config config.Config
	DB     *gorm.DB
)

func Md5ToString(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Sha256ToString(str string) string {
	data := []byte(str)
	has := sha256.Sum256(data)
	sha256str := fmt.Sprintf("%x", has)
	return sha256str
}

// func GetStructFieldValue(data interface{}, field string) {
// 	rs := reflect.ValueOf(data)
// 	fmt.Printf("%+v", rs)
// }
