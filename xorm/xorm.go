package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

type Reader struct {
	Id          int    `xorm:"int(11) notnull pk autoincr 'id'"`
	Name        string `xorm:"varchar(255) unique notnull 'name'"`
	Password    string `xorm:"varchar(255) notnull 'password'"`
	Sex         string `xorm:"varchar(255) notnull 'sex'"`
	BorrowedNum int    `xorm:"int(11) notnull 'borrowed_num'"`
}

func main() {

	//连接mysql
	engine, err := xorm.NewEngine(
		"mysql",
		"root:123456@tcp(192.168.2.33:3306)/library?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer engine.Close()

	//将日志写入到本地
	f, err := os.Create("1.log")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	results, err := engine.Query("select * from reader")

	for _, val := range results {
		fmt.Println(string(val["id"]))
		fmt.Println(string(val["name"]))
	}
	//插入数据
	{
		reader := Reader{Name: "test", Password: "test", Sex: "男", BorrowedNum: 0}
		flag, err := engine.Insert(reader)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(flag)
	}
	//删除数据
	{
		reader := Reader{Id: 5}
		flag, err := engine.Delete(reader)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(flag)
	}
	//修改数据
	{
		reader := Reader{Id: 6, Password: "1111"}
		flag, err := engine.ID(6).Update(reader)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(flag)
	}
	//	查询数据
	{
		var reader Reader
		_, err := engine.ID(14).Get(&reader)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(reader)
	}
	//查询所有的数据
	{
		readers := make([]Reader, 0)
		err := engine.Find(&readers)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(readers)
	}
	//条件查询where
	{
		readers := make([]Reader, 0)
		err := engine.Where("id> ? and borrowed_num > ?", 10, 3).Find(readers)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(readers)
	}
}
