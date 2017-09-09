package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := dbx.Open("mysql", "root:@/yiispace")

	// create a new query
	q := db.NewQuery("SELECT id, username FROM tbl_user LIMIT 10")

	// fetch all rows into a struct array
	var users []struct {
		ID, Username string
	}
	q.All(&users)

	// fetch a single row into a struct
	var user struct {
		ID, Username string
	}
	q.One(&user)

	fmt.Println(user)

	// fetch a single row into a string map
	data := dbx.NullStringMap{}
	q.One(data)

	/*
	// fetch row by row
	rows2, _ := q.Rows()
	for rows2.Next() {
		rows2.ScanStruct(&user)
		// rows.ScanMap(data)
		// rows.Scan(&id, &name)
	}
	*/
	fmt.Println("ok!")
}
