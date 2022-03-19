package model

import "github.com/jinzhu/gorm"

type Net struct {
	ID     int
	Command string
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (n Net) TableName() string {
	return "net"
}
func (p Product) TableName() string {
	return "product"
}



