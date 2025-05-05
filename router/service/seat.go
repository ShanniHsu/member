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
		exists, _ := rdb.Exists(context.Background(), fmt.Sprintf("seat:%s:status", seat)).Result()
		if exists != 0 {
			fmt.Println("Seats already initialized, skip.")
			continue
		}
		// 座位初始都是 available
		err := rdb.Set(context.Background(), fmt.Sprintf("seat:%s:status", seat), "available", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Seats initialized!")
}
