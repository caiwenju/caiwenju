package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DSN = "barry:barry@tcp(52.81.16.160:3306)/publish_dev?charset=utf8mb4&parseTime=True&loc=Local"
const DRIVER = "mysql"

func Db() *gorm.DB {
	db, err := gorm.Open(DRIVER, DSN)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	//var models = []interface{}{&model.Net{}, &model.Product{}}
	//db.AutoMigrate(models...)
	return db
}
