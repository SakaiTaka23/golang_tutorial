package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("astaxie1", 0777)
	err = os.Mkdir("astaxie2", 0777)
	err = os.MkdirAll("astaxie3/test1/test2", 0777)
	err = os.MkdirAll("astaxie4/test1/test2", 0777)
	err = os.Remove("astaxie2")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("astaxie4")
}