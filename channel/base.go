package channel

import "fmt"

/*
無緩衝通道的基本操作

建立一個無緩衝的 channel，名稱為 ch。
啟動一個 goroutine，該 goroutine 會向 ch 發送數字 100。
在 main 函數中接收該數字並打印出來。
*/
func TestOne() {
	var intChan = make(chan int)
	go func() {
		intChan <- 100
	}()
	x := <-intChan
	fmt.Println(x)
}
