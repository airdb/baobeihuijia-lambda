package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

func Connect(gdbc string) (db *gorm.DB, err error) {
	fmt.Println("gorm models connect")
	// conn, err = gorm.Open("mysql", "root:dean@tcp(msql.airdb.io:3306)/airdb?charset=utf8&parseTime=true&loc=Local")
	conn, err = gorm.Open("mysql", gdbc)

	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Article{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Comment{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Contact{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Login{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&User{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Summary{})
	conn.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&Image{})

	fmt.Println("models connect err:", err)
	return
}
