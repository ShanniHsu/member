package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"member/models"
	"member/pkg/argon2"
	"member/pkg/ecPay/proxy"
	"member/pkg/jwt"
	"member/pkg/uuid"
	"member/router/app/content/create-order"
	"member/router/app/content/get_user"
	"member/router/app/content/login"
	"member/router/app/content/register"
	"member/router/repository"
)

type User interface {
	Register(req *register.Request) (err error)
	Login(req *login.Request) (jwtToken string, err error)
	AuthBearerToken(token string) (user *models.User, err error)
	GetUserInfo(ctx *gin.Context) (resp *get_user.Response, err error)
	CreateOrder(req *create_order.Request) (body []byte, err error)
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
			var password string
			password, err = argon2.GenerateFromPassword(req.Password)
			if err != nil {
				err = errors.New("The password generate failed!")
				return
			}

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

func (s userService) Login(req *login.Request) (jwtToken string, err error) {
	resp, err := s.repo.UserRepository.GetUserByAccount(req.Account)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("Authentication failed!")
		return
	}
	match, err := argon2.ComparePasswordAndHash(req.Password, resp.Password)
	if err != nil {
		return
	}
	if !match {
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

func (s userService) GetUserInfo(ctx *gin.Context) (resp *get_user.Response, err error) {
	var user = new(models.User)
	ctxUser, exist := ctx.Get("user")
	if exist {
		user = ctxUser.(*models.User)
	}

	resp = &get_user.Response{
		Account:  user.Account,
		Nickname: user.Nickname,
		Status:   user.Status,
	}

	return
}

func (s userService) CreateOrder(req *create_order.Request) (body []byte, err error) {
	body, err = proxy.CreateOrder(req)
	if err != nil {
		err = errors.New("Create order failed")
		return
	}
	return
}
