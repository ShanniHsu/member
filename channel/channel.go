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

var cc = make(chan int)

func Ch() {
	go chHandle()
	// 讀取channel方式，要用for、v, ok := <-cc
	for {
		v, ok := <-cc
		if !ok {
			return
		}
		fmt.Println(v)
	}
}

func chHandle() {
	for i := 0; i < 10; i++ {
		cc <- i
	}
	close(cc)
}
