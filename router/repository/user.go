package repository

import (
	"gorm.io/gorm"
	"member/models"
	"member/pkg/storage"
)

type userRepository struct {
	DB *gorm.DB
}

func UerRepository() userRepository {
	db := storage.InitStorage.GetDBConnect()
	return userRepository{
		DB: db,
	}
}

func (r userRepository) GetUserByAccount(user models.User, account string) (resp models.User, err error) {
	err = r.DB.First(&user).Where("account = ?", account).Error
	return resp, err
}
