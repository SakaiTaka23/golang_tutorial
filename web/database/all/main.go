package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    //"time"
)

func main() {
	db, err := sql.Open("mysql", "root:abc123@/gotest?charset=utf8")
    checkErr(err)

    //データの挿入
    stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
    checkErr(err)

	var res sql.Result
	for i := 0; i < 10; i++ {
		res, err = stmt.Exec("astaxie", "研究開発部門", "2012-12-09")
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)
	    fmt.Println("LastInsertId:", id)
	}

    //データの更新
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("Update(id:1)", 1)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println("RowsAffected", affect)

	//データの削除
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)

    res, err = stmt.Exec(2)
    checkErr(err)

    affect, err = res.RowsAffected()
	checkErr(err)

    //データの検索
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    fmt.Println("RowsAffected", affect)
    db.Close()

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}