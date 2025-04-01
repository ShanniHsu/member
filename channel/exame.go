package channel

import (
	"fmt"
	"time"
)

/*
1. 單向通道
請寫一個函數 sendData，它接收一個「只能發送數據」的通道（chan<- int），並依次向通道發送 1, 2, 3, 4, 5。
然後，在 main 函數中建立一個通道，並使用 sendData 發送數據，再從通道讀取數據並打印出來。
*/

/*
2. 通道阻塞與非阻塞讀寫
請建立一個無緩衝 channel，然後：

在 main 函數內嘗試寫入數據，觀察是否會阻塞。

使用 select 語句實現 非阻塞讀取與寫入，當無數據可讀時，打印 "No data received"，當無法寫入時，打印 `"No space to send"。
*/

/*
3. 兩個通道的數據合併
請建立兩個 channel，分別在兩個 goroutine 中發送數據，然後在 main 函數中，使用 select 監聽兩個 channel，並將收到的數據打印出來。
*/
func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func Exam() {
	//1.
	var ch = make(chan int)
	go sendData(ch)
	for val := range ch {
		fmt.Println("sendData-val:", val)
	}

	//2.
	var ch1 = make(chan int)
	// var ch1 = make(chan int, 1) 如果有緩衝區，他會執行Send 10 to channel

	// 如果非緩衝區，就必須加上goroutine去接收
	//time.Sleep(1 * time.Second)
	//go func() {
	//	<-ch1
	//}()

	// channel 接收數字
	select {
	case ch1 <- 10:
		fmt.Println("Send 10 to channel")
	default:
		fmt.Println("No space to send")
	}

	//3.
	var ch2 = make(chan string)
	var ch3 = make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Send ch2"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "Send ch3"
	}()

	// 這邊不加上default是為了等待channel接收到數據後才執行，讓他保持阻塞(blocking)，直到channel傳來數據
	select {
	case msg := <-ch2:
		fmt.Println("Receive that ch2 send: ", msg)
	case msg := <-ch3:
		fmt.Println("Receive that ch3 send: ", msg)
	}
}
