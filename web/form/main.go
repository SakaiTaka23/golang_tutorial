package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //urlが渡すオプションを解析します。POSTに対してはレスポンスパケットのボディを解析します（request body）
	//注意：もしParseFormメソッドがコールされなければ、以下でフォームのデータを取得することができません。
	fmt.Println("\nmethod:", r.Method)
	fmt.Println(r.Form) //これらのデータはサーバのプリント情報に出力されます
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //ここでwに書き込まれたものがクライアントに出力されます。
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nmethod:", r.Method) //リクエストを取得するメソッド
	if r.Method == "GET" {

		// 乱数生成
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("html/login.gtpl")
		t.Execute(w, token)
	} else {
		// 多重送信防止
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//tokenの合法性を検証します。
		} else {
			//tokenが存在しなければエラーを出します。
		}

		// パースなしでも出力可 ただし最初の一つのみ
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))

		//ログインデータがリクエストされ、ログインのロジック判断が実行されます。
		// パースしないと反映されない
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		fmt.Println("r.Form", r.Form)

		// xss対策
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //サーバに出力されます。
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //クライアントに出力されます。

		// エスケープ
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		if err != nil {
			fmt.Println(err)
		}

		// バリデーションしてみる
		errors := validate(r.Form)
		fmt.Println(errors)
	}
}

// バリデーション
func validate(form url.Values) (errors []string) {
	const requiredErr = "%sは必須です。"
	const numErr = "%sは数字で入力してください。"
	const rangeErr = "%sは0～99の値で入力してください。"

	if len(form["username"][0]) == 0 {
		errors = append(errors, fmt.Sprintf(requiredErr, "ユーザ名"))
	}
	if len(form["password"][0]) == 0 {
		errors = append(errors, fmt.Sprintf(requiredErr, "パスワード"))
	}

	if form.Get("age") != "" {
		getint, err := strconv.Atoi(form.Get("age"))
		if err != nil {
			errors = append(errors, fmt.Sprintf(numErr, "年齢"))
		} else {
			if getint < 0 || getint > 99 {
				errors = append(errors, fmt.Sprintf(rangeErr, "年齢"))
			}
		}
	}

	slice := []string{"apple", "pear", "banane"}
	check := false

	for _, v := range slice {
		if v == form.Get("fruit") {
			check = true
			break
		}
	}
	if check == false {
		errors = append(errors, fmt.Sprintf("フルーツの選択が不正です"))
	}

	return
}

func main() {
	http.HandleFunc("/", sayhelloName)       //アクセスのルーティングを設定します
	http.HandleFunc("/login", login)         //アクセスのルーティングを設定します
	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
