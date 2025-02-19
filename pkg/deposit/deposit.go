package deposit

import (
	"fmt"
	"sync"
	"time"
)

var mu = new(sync.Mutex)

var balance int64 = 1000 // 原始存款餘額

func deposit(amount int64) {
	mu.Lock() //獲取鎖，確保一個goroutine修改balance
	defer mu.Unlock()
	balance += amount
}

func UseDeposit() {
	// 當這邊goroutine次數變多(這邊嘗試修改成1000000000)，會造成以下結果
	/* 1. CPU使用率暴增(許多goroutine進入等待鎖[Blocking on Mutex]的狀態，導致CPU使用率激增)
	   (1) goroutine會爭奪鎖，因為mu.Lock()會讓其他goroutine等待直到鎖釋放，這會"降低併發效能"。
	   (2) 在10億次迴圈下，Go Runtime需要頻繁地進行goroutine切換(context switching)，進一步影響效能。*/

	/* 2. 記憶體暴增(可能OOM，Out of Memory)
	   (1) 執行10億次時，可能會創建大量goroutine，如果機器的記憶體不夠，可能導致OOM崩潰。
	    Solution: 限制goroutine數量(如使用sync.WaitGroup或worker pool)*/

	/* 3. 計算時間過長
	   即使不發生OOM，因為單次mu.Lock()需要等待釋放後才能繼續，這會導致程式執行時間大幅拉長
	*/
	for i := 0; i < 1000000000; i++ {
		go deposit(10)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("balance: ", balance)
}
