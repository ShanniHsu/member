package channel

import "fmt"

var ch = make(chan int)

func Channel() {

	go handle()

	// 讀取channel方式，要用for range
	for v := range ch {
		fmt.Println(v)
	}
}

func handle() {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
