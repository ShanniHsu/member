package service

import (
	"context"
	"errors"
	"fmt"
	"member/pkg/storage"
	get_seats "member/router/app/content/get-seats"
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
	GetSeats() (resp []get_seats.Response, err error)
}

type seat struct {
	repo repository.Repo
}

func NewSeatService(repo repository.Repo) SeatService {
	return &seat{
		repo: repo,
	}
}

func (s *seat) GetSeats() (resp []get_seats.Response, err error) {
	value, err := s.repo.SeatRepository.ScanRedis(0, "seat:*:status", 0)
	if err != nil {
		err = errors.New("redis scan error")
		return
	}

	resp = make([]get_seats.Response, 0, len(value))
	for k, v := range value {
		item := get_seats.Response{
			Seat:   k,
			Status: v,
		}
		resp = append(resp, item)
	}
	return
}
