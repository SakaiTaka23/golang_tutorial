package gotest

import (
	"testing"
)

func TestDivision(t *testing.T) {

	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除算関数のテストが通りません") // もし予定されたものでなければエラーを発生させます。
	} else {
		t.Log("はじめのテストがパスしました") //記録したい情報を記録します
	}
}

func Test_Division_2(t *testing.T) {

	if i, e := Division(6, 1); e != nil { //try a unit test on function
		t.Error("Division did not work as expected.",e) // 予期したものでなければエラーを発生
	} else {
		t.Log("one test passed.", i) //記録したい情報を記録
	}
}
