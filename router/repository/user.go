package repository

import (
	"gorm.io/gorm"
	"member/models"
)

// method
type UserRepository interface {
	GetUserByAccount(account string) (resp *models.User, err error)
	Create(user *models.User) (err error)
	Update(user *models.User, newData map[string]interface{}) (err error)
}

// the request that is applied the method
type userRepository struct {
	DB *gorm.DB
}

// use userRepository to return UserRepository(interface)
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

// Mothod of UserRepository(interface)
func (r userRepository) GetUserByAccount(account string) (resp *models.User, err error) {
	err = r.DB.Where("account = ?", account).First(&resp).Error
	return resp, err
}

func (r userRepository) Create(user *models.User) (err error) {
	return r.DB.Create(&user).Error
}

func (r userRepository) Update(user *models.User, newData map[string]interface{}) (err error) {
	return r.DB.Model(&user).Updates(newData).Error
}
