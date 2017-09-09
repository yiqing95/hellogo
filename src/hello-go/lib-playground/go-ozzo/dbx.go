package main

import (
	"fmt"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := dbx.Open("mysql", "root:@localhost/yiispace")

	// create a new query
	q := db.NewQuery("SELECT id, name FROM user LIMIT 10")

	// fetch all rows into a struct array
	var users []struct {
		ID, Name string
	}
	q.All(&users)

	// fetch a single row into a struct
	var user struct {
		ID, Username, Name string
	}
	q.One(&user)

	fmt.Println(user)

	// fetch a single row into a string map
	data := dbx.NullStringMap{}
	q.One(data)

	q2 := db.NewQuery("UPDATE tbl_user SET status=14 WHERE id=66")
	result, err := q2.Execute()

	if err != nil {
		fmt.Print("OK la")
		fmt.Println(result)
	}

}
