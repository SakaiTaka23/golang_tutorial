package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// バッファを追加することにより複数のint数字を保存できるようにする
	// 送受信の時間が同じではないためどちらかが時間がかかったときに保存される
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		// 一度チャンネルを閉じると再度開くことはできない
		// しかし、チャンネルを閉めないと17行目はforの長さがわからない
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

}
