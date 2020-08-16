package main

import (
	"fmt"
	"golang_tutorial/blogsystem/controllers"
	"net/http"

	"github.com/astaxie/beego"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World<h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>about<h1>")
}

func main() {
	//ブログのトップページを表示
	beego.Router("/", &controllers.IndexController{})
	//ブログの詳細な情報を検索
	beego.Router("/view/:id([0-9]+)", &controllers.ViewController{})
	//ブログの文章を作成
	beego.Router("/new", &controllers.NewController{})
	//ブログの削除
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})
	//ブログの編集
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})
	
	http.ListenAndServe(":3000", nil)
}
