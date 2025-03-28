package channel

import (
	"fmt"
	"sync"
)

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

// https://blog.wu-boy.com/2022/05/read-data-from-channel-in-go/
// 練習使用兩個goroutine取值跑資料
func Foobar() {
	str := []byte("foobar")
	xch := make(chan byte, len(str))
	next := make(chan struct{}) // 空結構體通道（僅用來同步）
	wg := &sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < len(str); i++ {
		xch <- str[i]
	}
	close(xch)

	go func() {
		defer wg.Done()

		for {
			<-next
			v, ok := <-xch
			if !ok {
				close(next)
				return
			}
			fmt.Println("goroutine01: ", string(v))
			next <- struct{}{} // 用來通知，這個不會有數據
		}
	}()

	go func() {
		defer wg.Done()

		for {
			<-next
			v, ok := <-xch
			if !ok {
				close(next)
				return
			}
			fmt.Println("goroutine02: ", string(v))
			next <- struct{}{} // 用來通知，這個不會有數據
		}

	}()
	next <- struct{}{}
	wg.Wait()
}
