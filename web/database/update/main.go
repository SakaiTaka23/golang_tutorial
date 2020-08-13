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

	//データの更新
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("astaxieupdate", 5)
	checkErr(err)

	//更新があった列の数を返す
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
