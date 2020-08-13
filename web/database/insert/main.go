package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//接続
	db, err := sql.Open("mysql", "root:abc123@/gotest?charset=utf8")
	checkErr(err)

	//データの挿入
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研究開発部門", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("stmt",stmt)
	fmt.Println("res", res)
	fmt.Println("id", id)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
