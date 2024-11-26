package service

import (
	"errors"
	"gorm.io/gorm"
	"member/models"
	"member/pkg/jwt"
	"member/pkg/uuid"
	"member/router/app/content/get_user"
	"member/router/app/content/login"
	"member/router/app/content/register"
	"member/router/repository"
)

type User interface {
	Register(req *register.Request) (err error)
	Login(req *login.Request) (jwtToken string, err error)
	AuthBearerToken(token string) (user *models.User, err error)
	GetUserInfo(id int64) (resp *get_user.Response, err error)
}

type userService struct {
	repo repository.Repo
}

func NewUserService(repo repository.Repo) User {
	return &userService{
		repo: repo,
	}
}

func (s userService) Register(req *register.Request) (err error) {

	resp, err := s.repo.UserRepository.GetUserByAccount(req.Account)
	if resp.ID != 0 {
		err = errors.New("The account is existed!")
		return
	}

	//if the account is not existed
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
	}
	return
}

func (s userService) Login(req *login.Request) (jwtToken string, err error) {
	resp, err := s.repo.UserRepository.GetUserByAccount(req.Account)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("Authentication failed!")
		return
	}
	if resp.Password != req.Password {
		err = errors.New("Authentication failed!")
		return
	}

	// 給登入者身份識別
	token := uuid.GenerateUuid()

	newData := map[string]interface{}{
		"token": token,
	}

	err = s.repo.UserRepository.Update(resp, newData)
	if err != nil {
		err = errors.New("Update user failed!")
		return
	}

	jwtToken, err = jwt.GenerateJWT(token)
	if err != nil {
		return
	}
	return
}

func (s userService) AuthBearerToken(token string) (user *models.User, err error) {
	user, err = s.repo.UserRepository.GetUserByToken(token)
	if err != nil {
		err = errors.New("Token isn't found!")
		return
	}
	return
}

func (s userService) GetUserInfo(id int64) (resp *get_user.Response, err error) {
	var user *models.User
	user, err = s.repo.UserRepository.GetUserByID(id)
	if err != nil {
		err = errors.New("Get user failed!")
		return
	}

	resp = &get_user.Response{
		Account:  user.Account,
		Password: user.Password,
		Nickname: user.Nickname,
		Status:   user.Status,
	}

	return
}
