package deposit

import (
	"fmt"
	"time"
)

var balance int64 = 1000 // 原始存款餘額

func deposit(amount int64) {
	balance += amount
}

// 因為沒有使用鎖，多個gorutine同時讀取balance，導致寫入時發生競爭條件(race condition)，造成錯誤結果(這邊列印出來是10640而非預期結果11000)
func UseDeposit() {
	for i := 0; i < 1000; i++ {
		go deposit(10)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("balance: ", balance)
}
