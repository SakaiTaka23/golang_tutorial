package main

import (
	"fmt"
	"time"
)

var locales map[string]map[string]string

func main() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en
	cn := make(map[string]string, 10)
	cn["pea"] = "ピーナッツ"
	cn["bean"] = "枝豆"
	locales["ja-JP"] = cn
	lang := "ja-JP"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))

	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Tokyo"

	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))
	en["date_format"] = "%Y-%m-%d %H:%M:%S"
	cn["date_format"] = "%Y年 %m月%d日 %H時%M分%S秒"

	en["money"] = "USD %d"
	cn["money"] = "￥%d円"
	fmt.Println(moneyFormat(msg(lang, "moneyFormat"),100))
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}

func moneyFormat(fomate string, money int64) string {
	return fmt.Sprintf(fomate, money)
}
