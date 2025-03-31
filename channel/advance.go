package channel

import (
	"fmt"
	"time"
)

/*
使用 select 處理多個通道

建立兩個 channel，名稱為 ch1 和 ch2，它們傳遞 string。
啟動兩個 goroutine，分別在 1 秒 和 2 秒 後向 ch1 和 ch2 發送 "Hello from ch1" 和 "Hello from ch2"。
在 main 函數中使用 select 來監聽 ch1 和 ch2，並打印最先收到的訊息。
*/

func TestOneAdvance() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
