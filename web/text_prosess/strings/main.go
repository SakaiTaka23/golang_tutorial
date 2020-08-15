package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains 文字列が含まれているかをboolで変換
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	// Join 文字列結合
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	// Index 検索文字列のインデックスを返す なければ-1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	// 第４引数で置換回数指定 ０以下の場合全て
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))

}
