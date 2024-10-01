package newTicker

import (
	"fmt"
	"time"
)

func Ticker() {
	// 每2秒跑一次，會影響下面迴圈，決定每隔幾秒出現的值
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	// 同時開啟新的子執行緒 10秒後將true傳入done這變數(計時倒數10秒)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
