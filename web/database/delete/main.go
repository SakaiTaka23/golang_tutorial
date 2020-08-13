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


	//データの削除
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(12)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
