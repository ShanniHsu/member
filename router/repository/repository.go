package repository

import "member/pkg/storage"

type Repo struct {
	UserRepository           UserRepository
	RestaurantRepository     RestaurantRepository
	UserRestaurantRepository UserRestaurantRepository
}

func NewRepository() Repo {
	db := storage.InitStorage.GetDBConnect()
	redis := storage.InitStorage.GetRDBConnect()
	return Repo{
		UserRepository:           NewUserRepository(db, redis),
		RestaurantRepository:     NewRestaurantRepository(db),
		UserRestaurantRepository: NewUserRestaurantRepository(db),
	}
}
