package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// 順番通りには動くがロック、解除を繰り返すことにより
// goroutinesの恩恵を受けておらず、むしろシングルスレッド
// の時よりパフォーマンスが落ちている
func main() {
	// 使用するスレッドの数を指定
	// 大きいとパフォーマンスが上がるが、大きすぎると
	// スケジュールが難しくなり遅くなる
	runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}
