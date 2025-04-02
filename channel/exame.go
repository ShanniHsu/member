package channel

import (
	"fmt"
	"sync"
	"time"
)

// 基礎題
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

// 進階題
/*
4. 10 個 goroutine 同時計算，使用 channel 收集結果
請寫一個程式：

啟動 10 個 goroutine，每個 goroutine 計算 i * i (i 為 0~9)，並將結果發送到 channel。

main 函數收集 channel 中的結果，並打印出來。
*/

/*
5. 競態條件測試
請寫一段程式：

建立一個 int 變數 counter，啟動 5 個 goroutine，每個 goroutine 嘗試 同時 增加 counter 的值 1000 次。

在 main 函數中等待所有 goroutine 結束，然後打印 counter。

觀察結果是否為 5000，如果不是，請修正程式使其正確運行。
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

	//4.
	var ch4 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			sum := i * i
			ch4 <- sum
		}(i)
	}

	/*這邊不使用因為for range會一直等待讀取channel，直到channel被關閉，使用會deadlock
	for value := range ch4 {
		fmt.Println("value:", value)
	}*/

	for i := 0; i < 10; i++ {
		fmt.Println("第4題: ", <-ch4)
	}

	//5.
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for j := 0; j < 5; j++ {
		wg.Add(1) // 增加一個計數
		go func() {
			defer wg.Done() // 減少一個計數，表示goroutine完成
			for i := 0; i < 1000; i++ {
				mu.Lock()
				counter++ // 主要是有多個goroutine會同時修改counter，有可能會出現不對的數值，所以才需要加上鎖
				mu.Unlock()
			}
		}()
	}
	wg.Wait() // 阻塞主goroutine，直到計數歸零才繼續執行
	fmt.Println("counter:", counter)

}
