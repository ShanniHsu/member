package repository

import (
	"gorm.io/gorm"
	"member/models"
)

// method
type UserRepository interface {
	GetUserByAccount(account string) (user *models.User, err error)
	GetUserByID(id int64) (resp *models.User, err error)
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
func (r userRepository) GetUserByAccount(account string) (user *models.User, err error) {
	err = r.DB.Where("account = ?", account).First(&user).Error
	return
}

func (r userRepository) GetUserByID(id int64) (user *models.User, err error) {
	user = new(models.User)
	err = r.DB.Where("id = ?", id).First(&user).Error
	return
}

func (r userRepository) Create(user *models.User) (err error) {
	return r.DB.Create(&user).Error
}

func (r userRepository) Update(user *models.User, newData map[string]interface{}) (err error) {
	return r.DB.Model(&user).Updates(newData).Error
}
