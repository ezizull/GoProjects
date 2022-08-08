package config

import (
	_ "github.com/go-sql-driver/mysql" // "mysql"
	"github.com/jinzhu/gorm"
)

// var exported
var (
	DBConn *gorm.DB
)


// MySQLConnect exported
func MySQLConnect(){
	db, err := gorm.Open("mysql", "root:@/goproject?charset=utf8&parseTime=True&loc=Local")
	if err != nil { panic(err) }

	DBConn = db
}

// GetDB exported
func GetDB() *gorm.DB{
	return DBConn
}