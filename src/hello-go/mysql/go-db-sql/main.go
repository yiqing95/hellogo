package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, err := sql.Open("mysql",
		// 竟然不检测连接参数的正确性 这是一个懒对象！
		"yiqing:yiqing@tcp(127.0.0.1:3306)/test")
	// "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 养成好习惯
	// 实际上 db 需要全局共享 也不需要频繁的 open 和 close 最好作为依赖传递给 function 或者对象

	err = db.Ping()
	if err != nil {
		log.Fatalln("failed to connect db :", err)
	}
	log.Println("connected successfully !")

	query1(db)
	preparingQuery(db)
	singleRowQuery(db)
	preparingQueryRow(db)

	// -----------------------------

	// insertUser(db)
	timeTest()
}

func query1(db *sql.DB) {
	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from user where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func preparingQuery(db *sql.DB) {
	stmt, err := db.Prepare("select id, name from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// ...
		user := User{}
		rows.Scan(&user.Id, &user.Name)
		fmt.Println(user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func singleRowQuery(db *sql.DB) {
	var name string
	// Errors from the query are deferred until Scan() is called, and then are returned from that
	// 错误被延迟到scan方法的调用; queryXxx 后已经有error了！不过是私有的 只有在Scan后才被公开暴露
	err := db.QueryRow("select name from user where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func preparingQueryRow(db *sql.DB) {
	stmt, err := db.Prepare("select name from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	var name string
	err = stmt.QueryRow(1).Scan(&name)
	if err != nil {
		// 此处处理参考 handleError方法 在未找到记录时候 也是会报错的 应用一般不认为其是致命（fatal）错误
		log.Fatal(err)
	}
	fmt.Println(name)
}

/**
*  better way  than above function
 */
func handleError(db *sql.DB) {
	stmt, err := db.Prepare("select name from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	var name string
	err = stmt.QueryRow(1).Scan(&name)
	if err != nil {
		// 注意 只有在QueryRow 中会出现这种特殊情况 其他地方如果出现了 就意味着你犯错误了！
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}

func insertUser(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(name,email,created_at) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("qing2", "yiqing2@china.com", time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d , affected = %d\n", lastId, rowCnt)

}

func timeTest() {
	// @see http://blog.csdn.net/u012210379/article/details/44747311
	//获取本地location
	toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)
}
