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
	chx := make(chan byte, len(str))
	next := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(str); i++ {
		chx <- str[i]
	}
	close(chx)

	go func() {
		defer wg.Done()

		for {
			stop, ok := <-next
			if !ok {
				return
			}
			v, ok := <-chx
			if !ok {
				close(next)
				return
			}
			fmt.Println("goroutine01: ", string(v))
			next <- stop
		}

	}()

	go func() {

		defer wg.Done()

		for {
			stop, ok := <-next
			if !ok {
				return
			}
			v, ok := <-chx
			if !ok {
				close(next)
				return
			}
			fmt.Println("goroutine02: ", string(v))
			next <- stop
		}

	}()
	next <- struct{}{}
	wg.Wait()
}
