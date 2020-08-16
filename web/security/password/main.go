package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	//import "crypto/sha256"
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))

	//import "crypto/sha1"
	h = sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))

	//import "crypto/md5"
	h = md5.New()
	io.WriteString(h, "暗号化する必要のあるパスワード")
	fmt.Printf("%x", h.Sum(nil))
}
