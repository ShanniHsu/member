package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"member/models"
	get_restaurants "member/router/app/content/get-restaurants"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"member/router/repository"
)

type Restaurant interface {
	GetRestaurants() (restaurants *[]models.Restaurant, err error)
	GetRestaurantList(ctx *gin.Context, req *get_restaurants.Request) (restaurants *[]models.Restaurant, err error)
}

type restaurantService struct {
	repo repository.Repo
}

func NewRestaurantService(repo repository.Repo) Restaurant {
	return restaurantService{
		repo: repo,
	}
}

func (s restaurantService) GetRestaurants() (restaurants *[]models.Restaurant, err error) {
	restaurants, err = s.repo.RestaurantRepository.GetRestaurants()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("Restaurants isn't found!")
		return
	}
	return
}

func (s restaurantService) GetRestaurantList(ctx *gin.Context, req *get_restaurants.Request) (restaurants *[]models.Restaurant, err error) {
	restaurantAll, err := s.repo.RestaurantRepository.GetRestaurantFilter(req)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("Restaurants isn't found!")
		return
	}

	var user = new(models.User)
	ctxUser, exist := ctx.Get("user")
	if exist {
		user = ctxUser.(*models.User)
	}

	var parameter = new(get_user_restaurants.Request)
	newRestaurantList := *restaurantAll
	parameter.Type = req.Type
	parameter.Page = "1"
	parameter.PageSize = "500"
	userRestaurant, _ := s.repo.UserRestaurantRepository.GetUserRestaurantFilter(parameter, user.ID)

	if userRestaurant.TotalCount > 0 {
		for i := 0; i < len(userRestaurant.List); i++ {
			for key, value := range newRestaurantList {
				if value.ID == userRestaurant.List[i].RestaurantID {
					newRestaurantList = slices.Delete(newRestaurantList, key, key+1)
					break
				}
			}
		}
	}
	restaurants = &newRestaurantList
	return
}
