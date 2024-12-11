package service

import (
	"errors"
	"gorm.io/gorm"
	"member/models"
	"member/pkg/argon2"
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
	if resp.ID != 0 {
		err = errors.New("The account is existed!")
		return
	}

	password, err := argon2.GenerateFromPassword(req.Password)
	if err != nil {
		err = errors.New("The password generate failed!")
		return
	}

	//if the account is not existed
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var user = new(models.User)
			user = &models.User{
				Account:  req.Account,
				Password: password,
				Nickname: req.Account,
			}
			err = s.repo.UserRepository.Create(user)
			if err != nil {
				return err
			}
		}
	}
	return
}
