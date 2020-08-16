package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	//import "crypto/md5"
	//ユーザ名をabc、パスワードを123456とします
	h := md5.New()
	io.WriteString(h, "暗号化が必要なパスワード")

	//pwmd5はe10adc3949ba59abbe56e057f20f883eです。
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	println(pwmd5)

	//saltを２つ指定します： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+ユーザ名+salt2+MD5を連結します。
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)

}
