package repository

import "member/pkg/storage"

type Repo struct {
	UserRepository UserRepository
}

func NewRepository() Repo {
	db := storage.InitStorage.GetDBConnect()
	return Repo{
		UserRepository: NewUserRepository(db),
	}
}
