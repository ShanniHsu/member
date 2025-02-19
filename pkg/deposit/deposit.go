package deposit

import (
	"fmt"
	"sync"
	"time"
)

var mu = new(sync.Mutex)

var balance int64 = 1000 // 原始存款餘額

const workerCount = 100 // 限制最多100個goroutine

func deposit(amount int64, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range jobs {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}
}

func UseDeposit() {
	t := time.Now()
	// 建立一個channel，作為工作佇列，存放要執行的工作
	jobs := make(chan int, 1000) // 任務佇列
	var wg sync.WaitGroup

	// 啟用固定數量的goroutine (100個deposit)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go deposit(10, jobs, &wg)
	}

	// 送10億個存款任務
	for i := 0; i < 1000000000; i++ {
		jobs <- 0 //代表有接收
	}
	close(jobs) // 所有工作都放入後關閉通道

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("balance: ", balance)
	fmt.Println("Spend time:", time.Since(t))
}
