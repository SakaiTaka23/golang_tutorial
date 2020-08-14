package main

import (
	"fmt"
	"regexp"
)

// IsIP ipアドレスを正規表現によって判定
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	ipaddr := "192.168.12.1"
	fmt.Println(IsIP(ipaddr))
}
