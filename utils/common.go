package common

import (
	"fmt"
	// "fmt"
	"red-east/config"
	"reflect"

	// "red-east/dao/database"
	"red-east/utils/logging"

	"github.com/jinzhu/gorm"
)

var (
	Logger logging.NLogger
	Config config.Config
	DB     *gorm.DB
)

func Md5ToString() string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func GetStructFieldValue(data interface{}, field string) {
	rs := reflect.ValueOf(data)
	fmt.Printf("%+v", rs)
}
