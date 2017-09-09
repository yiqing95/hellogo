package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	// engine, err = xorm.NewEngine("mysql", "yiqing:yiqing@tcp/yiispace?charset=utf8")
	engine, err = xorm.NewEngine("mysql", "yiqing:yiqing@tcp(localhost:3306)/yiispace?charset=utf8")

	err = engine.Ping()
	if err != nil {
		fmt.Println("连不上数据库啦！")
		fmt.Println(err)
	} else {
		fmt.Println("连接成功！")
	}

	// 记录日志
	engine.ShowSQL(true)

	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
}
