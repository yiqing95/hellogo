package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/yii_space")
	if err != nil {
		log.Fatalln("Fatal err when connect mysql :", err)
	}

	fmt.Println("connectted success !", db)
	// select 语句中的 字段如果是* 在scan时 必须写所有字段？ 不然填充不上！
	rows, err := db.Query("select id, username , email from user")
	if err != nil {
		fmt.Println(err)
		panic("query user failed !")
	}
	// Get column names
	columns, err := rows.Columns()
	fmt.Println(columns)

	defer rows.Close()
	for rows.Next() {
		var id int
		var username string
		var email string

		err = rows.Scan(&id, &username, &email)

		// fmt.Printf("user:  %s \n ", email)
		fmt.Printf("user: %d %s %s \n ", id, username, email)
	}
	fmt.Println(rows.Err())

	queryRow(db)
	addUserData(db)
}

func queryRow(db *sql.DB) {
	var username string
	row := db.QueryRow("select username from user where id=?", 1)
	err := row.Scan(&username)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("no result with id: %d ", 1)
	case err != nil:
		log.Fatalln(err)
	default:
		fmt.Println("username : ", username)
	}
}

func addUserData(db *sql.DB) {
	rslt, err := db.Exec("insert into user_data(user_id,attr, val) values(?,?,?)",
		2, "test_attr", "test_val")

	if err != nil {
		log.Fatal(err)
	}
	ra, _ := rslt.RowsAffected()
	id, err := rslt.LastInsertId()
	if err != nil {
		log.Println("result err :", err)
	}
	log.Println("Rows affected ", ra, " last inserted id :", id)

}
