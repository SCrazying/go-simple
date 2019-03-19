package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {

	db, err := gorm.Open("mysql", "root:123456@(192.168.0.16:3306)/library?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败 ", err)
		panic(err)
	} else {
		fmt.Println("连接db成功")
		Db = db
	}

}
