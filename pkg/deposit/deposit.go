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
	for i := 0; i < 1000; i++ {
		go deposit(10)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("balance: ", balance)
}
