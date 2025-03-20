package repository

import (
	"gorm.io/gorm"
	"member/models"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"strconv"
)

type UserRestaurantRepository interface {
	GetUserRestaurantFilter(parameter *get_user_restaurants.Request, userID int64) (resp *get_user_restaurants.Response, err error)
	DeleteByID(id int64) (err error)
	Create(userRestaurant *models.UserRestaurant) (err error)
}

type userRestaurantRepository struct {
	DB *gorm.DB
}

func NewUserRestaurantRepository(db *gorm.DB) UserRestaurantRepository {
	return userRestaurantRepository{
		DB: db,
	}
}

func (r userRestaurantRepository) GetUserRestaurantFilter(parameter *get_user_restaurants.Request, userID int64) (resp *get_user_restaurants.Response, err error) {
	page, _ := strconv.Atoi(parameter.Page)
	pageSize, _ := strconv.Atoi(parameter.PageSize)
	offset := (page - 1) * pageSize

	resp = &get_user_restaurants.Response{} // 指標且初始化(確保有可寫入的記憶體 == new(get_user_restaurants.Response)
	query := r.DB.Model(&models.Restaurant{}).
		Select("`user_restaurants`.id, `restaurants`.name, `restaurants`.address, `restaurants`.type").
		Joins("JOIN user_restaurants ON restaurants.id = user_restaurants.restaurant_id").
		Where("user_restaurants.user_id = ?", userID)

	if parameter.ID != 0 {
		query = query.Where("user_restaurants.id = ?", parameter.ID)
	}

	if parameter.Type != 0 {
		query = query.Where("restaurants.type = ?", parameter.Type)
	}

	if parameter.Name != "" {
		query = query.Where("restaurants.name = ?", parameter.Name)
	}

	if parameter.Address != "" {
		query = query.Where("restaurants.address = ?", parameter.Address)
	}

	query = query.Offset(offset).Limit(pageSize).Find(&resp.List)
	query = query.Count(&resp.TotalCount)
	return
}

func (r userRestaurantRepository) DeleteByID(id int64) (err error) {
	return r.DB.Delete(&models.UserRestaurant{ID: id}).Error
}

func (r userRestaurantRepository) Create(userRestaurant *models.UserRestaurant) (err error) {
	return r.DB.Create(userRestaurant).Error
}
