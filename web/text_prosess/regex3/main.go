package main

import (
    "fmt"
    "regexp"
)

func main() {
    a := "I am learning Go language"

    re, _ := regexp.Compile("[a-z]{2,4}")

    //正規表現にマッチする最初のものを探し出す
    one := re.Find([]byte(a))
    fmt.Println("Find:", string(one))

    //正規表現にマッチするすべてのsliceを探し出す。nが0よりも小さかった場合はすべてのマッチする文字列を返します。さもなければ指定した長さが返ります。
    all := re.FindAll([]byte(a), -1)
    fmt.Println("FindAll", all)

    //条件にマッチするindexの位置を探し出す。開始位置と終了位置。
    index := re.FindIndex([]byte(a))
    fmt.Println("FindIndex", index)

    //条件にマッチするすべてのindexの位置を探し出す、nは同上
    allindex := re.FindAllIndex([]byte(a), -1)
    fmt.Println("FindAllIndex", allindex)

    re2, _ := regexp.Compile("am(.*)lang(.*)")

    //Submatchを探し出し、配列を返します。はじめの要素はマッチしたすべての要素です。２つ目の要素ははじめの()の中で、３つ目は２つ目の()の中です。
    //以下の出力でははじめの要素は"am learning Go language"です。
    //２つ目の要素は" learning Go "です。空白を含んで出力することに注意してください。
    //３つ目の要素は"uage"です。
    submatch := re2.FindSubmatch([]byte(a))
    fmt.Println("FindSubmatch", submatch)
    for _, v := range submatch {
        fmt.Println(string(v))
    }

    //定義と上のFindIndexは同じです。
    submatchindex := re2.FindSubmatchIndex([]byte(a))
    fmt.Println(submatchindex)

    //FindAllSubmatchは条件にマッチするすべてのサブマッチを探し出します。
    submatchall := re2.FindAllSubmatch([]byte(a), -1)
    fmt.Println(submatchall)

    //FindAllSubmatchIndexはすべてのサブマッチのindexを探し出します。
    submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
    fmt.Println(submatchallindex)
}