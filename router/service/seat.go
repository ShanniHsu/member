package service

import (
	"context"
	"errors"
	"fmt"
	"member/pkg/storage"
	"member/router/repository"
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

type SeatService interface {
	GetSeats() (err error)
}

type seat struct {
	repo repository.Repo
}

func NewSeatService(repo repository.Repo) SeatService {
	return &seat{
		repo: repo,
	}
}

func (s *seat) GetSeats() (err error) {
	err = s.repo.SeatRepository.ScanRedis(0, "seat:*:status", 0)
	if err != nil {
		err = errors.New("redis scan error")
		return
	}
	return
}
