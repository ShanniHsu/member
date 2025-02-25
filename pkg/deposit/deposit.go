package deposit

import (
	"fmt"
	"sync"
	"time"
)

var mu = new(sync.Mutex)

var balance int = 1000 // 原始存款餘額

const workerCount = 100 // 限制最多100個goroutine

const (
	numWorkers = 100       // worker數量
	numJobs    = 100000000 // 總共10億個任務
	batchSize  = 1000      // 1個worker一次處理1000個jobs
)

/* 優化:
   1. 使用go func()來生產任務，避免阻塞
   2. 批量處理(batchSize = 1000)，減少Channel I/O開銷，提高效率
   3. 減少worker間競爭，提升CPU/GPU使用率

   結果:
   1. 比原始方案快
   2. 減少Channel傳輸開銷
   3. 降低goroutine切換的成本
   4. 更適合大規模數據處理 (如大數據計算、爬蟲、批量API請求等)
*/

// 負責批量處理數據
func worker(jobs <-chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for batch := range jobs {
		for _, amount := range batch {
			mu.Lock()
			balance += amount
			mu.Unlock()
		}
		//time.Sleep(10 * time.Millisecond) // 模擬運算延遲
	}
}

func UseWorker() {
	t := time.Now()
	jobs := make(chan []int, numWorkers) // 使用Channel批量傳遞數據
	var wg sync.WaitGroup

	// 啟動Worker
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	// 生產10億個任務 (用goroutine避免阻塞)
	go func() {
		batch := []int{} //批量儲存
		for i := 0; i < numJobs; i++ {
			batch = append(batch, 10) // 假設任務是存入10
			if len(batch) == batchSize {
				jobs <- batch   // 傳送一批
				batch = []int{} // 重置批次
			}
		}

		// 確保剩下的也能傳送
		if len(batch) > 0 {
			jobs <- batch
		}

		// 所有工作分配完畢
		close(jobs)
	}()

	wg.Wait() // 等待所有worker完成
	fmt.Println("balance:", balance)
	fmt.Println("所有工作完成")
	fmt.Println("Spend time:", time.Since(t))
}
