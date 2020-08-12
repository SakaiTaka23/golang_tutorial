package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// チャンネルを使用するときは関数は読むもしくは書き込む
	// のいずれかに特化させる →引数を設定することにより明示
	ch := make(chan int)
	wg.Add(2)

	// reciever
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// sender
	go func(ch chan<- int) {
		i := 42
		ch <- i
		i = 27
		wg.Done()
	}(ch)

	wg.Wait()
}
