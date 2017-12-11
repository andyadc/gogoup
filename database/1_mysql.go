package main

/*
sql.Open()函数用来打开一个注册过的数据库驱动，
go-sql-driver中注册了mysql这个数据库驱动，第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息。
它支持如下格式：
	user@unix(/path/to/socket)/dbname?charset=utf8
	user:password@tcp(localhost:5555)/dbname?charset=utf8
	user:password@/dbname
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。

db.Query()函数用来直接执行Sql返回Rows结果。

stmt.Exec()函数用来执行stmt准备好的SQL语句

*/

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:andyadc@/idea?charset=utf8")
	checkErr(err)
	fmt.Println(db)

	// insert
	insertSql := `INSERT gogoup SET name = ? , age = ?`
	stmt, err := db.Prepare(insertSql)
	checkErr(err)

	result, err := stmt.Exec(`adc`, 12)
	checkErr(err)

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// update
	updateSql := `UPDATE gogoup SET name = ? WHERE id = ?`
	ustmt, err := db.Prepare(updateSql)
	checkErr(err)

	uresult, err := ustmt.Exec(`adc1`, 2)

	fmt.Println(uresult.RowsAffected())

	// query
	rows, err := db.Query("SELECT * FROM gogoup")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		checkErr(err)

		fmt.Printf("id: %d, name: %s, age:%d \n", id, name, age)
	}

	// delete
	deleteSql := `DELETE FROM gogoup WHERE id = ?`
	dstmt, err := db.Prepare(deleteSql)
	checkErr(err)

	dres, err := dstmt.Exec(10)
	checkErr(err)

	fmt.Println(dres.RowsAffected())

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
