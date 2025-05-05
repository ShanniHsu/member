package service

import (
	"context"
	"fmt"
	"member/pkg/storage"
)

func InitSeats() {
	rdb := storage.InitStorage.GetRDBConnect()
	seats := []string{"A1", "A2", "A3", "A4", "B1", "B2"}

	for _, seat := range seats {
		// 座位初始都是 available
		err := rdb.Set(context.Background(), fmt.Sprintf("seat:%s:status", seat), "available", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Seats initialized!")
}
