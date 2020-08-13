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

	//データの検索
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username, department, created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
		fmt.Printf("\n")
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
