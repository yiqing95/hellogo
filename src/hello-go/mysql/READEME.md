- [go-database-sql](http://go-database-sql.org/index.html)

-  处理错误

~~~go
    
    // 尽管可以但 不是一个好的方法
    rows, err := db.Query("SELECT someval FROM sometable")
    // err contains:
    // ERROR 1045 (28000): Access denied for user 'foo'@'::1' (using password: NO)
    if strings.Contains(err.Error(), "Access denied") {
        // Handle the permission-denied error
    }


    // 最好类型判断：
    if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
         //  代码坏味道
        if driverErr.Number == 1045 {
            // Handle the permission-denied error
        }
    }

   // 修改为
   if driverErr, ok := err.(*mysql.MySQLError); ok {
	if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
            // Handle the permission-denied error
        }
    }
~~~

-  空值  
空值在db中scan到 语言中变成两种可能类型 在强类型中实际是不允许的  这个问题在json传输中也存在

比如 int 字段 允许为空  那么最后你得到的值的类型 可能为null或者为int ！

~~~go
for rows.Next() {
	var s sql.NullString
	err := rows.Scan(&s)
	// check err
	if s.Valid {
	   // use s.String
	} else {
	   // NULL value
	}
}
~~~
上面的处理方法实际不是太好  因为没有 sql.NullInt64 ...

~~~go

    rows, err := db.Query(`
        SELECT
            name,
            COALESCE(other_field, '') as other_field
        WHERE id = ?
    `, 42)

    for rows.Next() {
        err := rows.Scan(&name, &otherField)
        // ..
        // If `other_field` was NULL, `otherField` is now an empty string. This works with other data types as well.
    }
~~~

比较好的 就是用sql语句提前转换一下可能为null的记录 到一个比较合理的值（比如 “零”值）
这也启示了我们在设计db字段时尽量不要允许其为null
