package channel

import "fmt"

/*
1. 單向通道
請寫一個函數 sendData，它接收一個「只能發送數據」的通道（chan<- int），並依次向通道發送 1, 2, 3, 4, 5。
然後，在 main 函數中建立一個通道，並使用 sendData 發送數據，再從通道讀取數據並打印出來。
*/
func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func Exam() {
	var ch = make(chan int)
	go sendData(ch)
	for val := range ch {
		fmt.Println("sendData-val:", val)
	}
}
