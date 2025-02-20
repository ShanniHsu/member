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
	// 但這邊還是不夠有效率
	/* 1. 太多任務同時塞入channel
	   Go Channel雖然很強大，但jobs是個有界通道(這邊大小是1000)，當jobs被填滿時，主goroutine會阻塞，導致執行效率下降。
	   Solution: 使用sync.WaitGroup & goroutine生產任務，讓生產與消費同時進行。
	   2. 不需要10億個worker
	   即使workerCount = 100，處理10億次依舊很慢，因為我們的worker仍然是一個一個處理。
	   Solution: 增加批量處理 (Batch Processing)，一次處理多個數據，提高吞吐量。
	*/
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
