package util

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:abc@/pay_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("init DB error", err)
		ErrorLog.Errorf("Open Mysql ERROR :%v\n", err)
	}
}
