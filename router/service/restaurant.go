package service

import (
	"errors"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"member/models"
	"member/pkg/message"
	get_restaurants "member/router/app/content/get-restaurants"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"member/router/repository"
)

type Restaurant interface {
	GetRestaurants() (restaurants []models.Restaurant, err error)
	GetRestaurantList(user *models.User, req *get_restaurants.Request) (restaurants []models.Restaurant, err error)
}

type restaurantService struct {
	repo repository.Repo
}

func NewRestaurantService(repo repository.Repo) Restaurant {
	return restaurantService{
		repo: repo,
	}
}

func (s restaurantService) GetRestaurants() (restaurants []models.Restaurant, err error) {
	restaurants, err = s.repo.RestaurantRepository.GetRestaurants()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New(message.GET_DATA_FAILED)
		return
	}
	return
}

func (s restaurantService) GetRestaurantList(user *models.User, req *get_restaurants.Request) (restaurants []models.Restaurant, err error) {
	restaurantAll, err := s.repo.RestaurantRepository.GetRestaurantFilter(req)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New(message.GET_DATA_FAILED)
		return
	}

	parameter := &get_user_restaurants.Request{
		Type:     req.Type,
		Page:     "1",
		PageSize: "500",
	}

	userRestaurant, _ := s.repo.UserRestaurantRepository.GetUserRestaurantFilter(parameter, user.ID)

	if userRestaurant.TotalCount > 0 {
		for i := 0; i < len(userRestaurant.List); i++ {
			for key, value := range restaurantAll {
				if value.ID == userRestaurant.List[i].RestaurantID {
					// 後續要優化，盡可能避免使用slice.Delete這類昂貴操作
					restaurantAll = slices.Delete(restaurantAll, key, key+1)
					break
				}
			}
		}
	}
	restaurants = restaurantAll
	return
}
