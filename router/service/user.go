package service

import (
	"errors"
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

	resp, err := s.repo.UserRepository.GetUserByAccount(req.Account)
	fmt.Println(resp.ID)
	if resp.ID != 0 {
		err = errors.New("The account is existed!")
		return
	}
	//if the account is not existed
	if resp.ID == 0 {
		var user = new(models.User)
		user = &models.User{
			Account:  req.Account,
			Password: req.Password,
			Nickname: req.Account,
		}
		err = s.repo.UserRepository.Create(user)
		if err != nil {
			return err
		}
	}
	return
}
