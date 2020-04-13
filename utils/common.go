package common

import (
	// "fmt"
	// "reflect"
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

// func GetStructField(paramStruct interface{}, key string) (conf interface{}, err error) {
// 	var value interface{}
// 	value = nil
// 	object := reflect.ValueOf(paramStruct)
// 	objectType := object.Type()
// 	for i := 0; i < object.NumField(); i++ {
// 		fieldName := objectType.Field(i)
// 		// fmt.Println(fieldName.Name)
// 		// fmt.Printf("%+v", fieldName)
// 		if fieldName.Name == key {
// 			value = object.Field(i)
// 		}
// 	}
// 	// fmt.Println(objecType.Field(0).Name)
// 	// fmt.Println(paramStruct)
// 	return value, err
// }
