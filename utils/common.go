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

func GetStructFieldValue(data interface{}, field string) {
	rs := reflect.ValueOf(data)
	fmt.Printf("%+v", rs)
}
