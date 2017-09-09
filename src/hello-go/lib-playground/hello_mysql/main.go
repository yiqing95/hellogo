package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)


func main()  {
	db, err := sql.Open("mysql", "yiqing:yiqing@/yii_space")
	if err == nil{
		fmt.Println(err)
	}

	errPing := db.Ping()
	fmt.Println(errPing)

	fmt.Println(db.Stats())
	db.Close()
}
