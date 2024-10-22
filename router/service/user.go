package service

import (
	"fmt"
	"member/models"
	"member/router/app/content/register"
	"member/router/repository"
)

type User interface {
	Register(req register.Request) (err error)
}

type userService struct {
	repo repository.Repo
}

func NewUserService(repo repository.Repo) User {
	return &userService{
		repo: repo,
	}
}

func (s userService) Register(req register.Request) (err error) {
	fmt.Println("req: ", req)
	resp, err := s.repo.UserRepository.GetUserByAccount(models.User{}, req.Account)
	fmt.Println("resp: ", resp, ";", "err: ", err)
	return
}
