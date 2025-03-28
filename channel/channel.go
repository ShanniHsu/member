package channel

import "fmt"

func Channel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 讀取channel方式，要用for range
	for v := range ch {
		fmt.Println(v)
	}
}
