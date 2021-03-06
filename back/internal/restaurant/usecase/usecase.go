package usecase

import (
	"context"
	"hurrles/config"
	"hurrles/internal/models"
	"hurrles/internal/restaurant/repository"
	"net/http"
	"time"
)

type IRestaurantUsecase interface {
	RestaurantsGet(context.Context) (models.RestaurantList, int, error)
	RestaurantIdPlacesGet(context.Context, models.PlaceParameters) (models.PlaceList, int, error)
	RestaurantIdAllPlacesGet(context.Context, uint64) (models.PlaceList, int, error)
	RestaurantIdMenuGet(context.Context, uint64) (models.DishList, int, error)
	RestaurantFavoritePost(context.Context, uint64) (int, error)
	RestaurantFavoriteGet(context.Context) (models.RestaurantList, int, error)
	RestaurantFavoriteDelete(context.Context, uint64) (int, error)
	RestaurantSearchPost(context.Context, string) (models.RestaurantList, int, error)
}

type restaurantUsecase struct {
	RestaurantPostgresRepository repository.IRestaurantRepository
	Timeout                      time.Duration
}

func NewRestaurantUsecase(rr repository.IRestaurantRepository, timeout time.Duration) IRestaurantUsecase {
	return &restaurantUsecase{
		RestaurantPostgresRepository: rr,
		Timeout:                      timeout,
	}
}

func (ru *restaurantUsecase) RestaurantsGet(ctx context.Context) (models.RestaurantList, int, error) {
	restaurants, err := ru.RestaurantPostgresRepository.GetRestaurants(ctx)
	if err != nil {
		return models.RestaurantList{}, http.StatusInternalServerError, err
	}
	return restaurants, http.StatusOK, nil
}

func (ru *restaurantUsecase) RestaurantIdPlacesGet(ctx context.Context, placeParams models.PlaceParameters) (models.PlaceList, int, error) {
	bookedPlaces, err := ru.RestaurantPostgresRepository.GetRestaurantBookedPlaces(ctx, placeParams)
	if err != nil {
		return models.PlaceList{}, http.StatusInternalServerError, err
	}
	places, err := ru.RestaurantPostgresRepository.GetRestaurantPlacesByFloor(ctx, placeParams)
	if err != nil {
		return models.PlaceList{}, http.StatusInternalServerError, err
	}
	for i := range places {
		for j := range bookedPlaces {
			if bookedPlaces[j].Id == places[i].Id {
				places[i].IsBooked = true
			}
		}
	}
	return places, http.StatusOK, nil
}

func (ru *restaurantUsecase) RestaurantIdAllPlacesGet(ctx context.Context, restaurantId uint64) (models.PlaceList, int, error) {
	places, err := ru.RestaurantPostgresRepository.GetRestaurantAllPlaces(ctx, restaurantId)
	if err != nil {
		return models.PlaceList{}, http.StatusInternalServerError, err
	}

	return places, http.StatusOK, nil
}

func (ru *restaurantUsecase) RestaurantIdMenuGet(ctx context.Context, id uint64) (models.DishList, int, error) {
	dishes, err := ru.RestaurantPostgresRepository.GetRestaurantMenuById(ctx, id)
	if err != nil {
		return models.DishList{}, http.StatusInternalServerError, err
	}
	return dishes, http.StatusOK, nil
}

func (ru *restaurantUsecase) RestaurantFavoritePost(ctx context.Context, restaurantId uint64) (int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return http.StatusNotFound, nil
	}

	err := ru.RestaurantPostgresRepository.CreateFavoriteRestaurant(ctx, curUser.Id, restaurantId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (ru *restaurantUsecase) RestaurantFavoriteGet(ctx context.Context) (models.RestaurantList, int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return models.RestaurantList{}, http.StatusNotFound, nil
	}

	favorites, err := ru.RestaurantPostgresRepository.GetRestaurantFavorite(ctx, curUser.Id)
	if err != nil {
		return models.RestaurantList{}, http.StatusInternalServerError, err
	}
	return favorites, http.StatusOK, nil
}

func (ru *restaurantUsecase) RestaurantFavoriteDelete(ctx context.Context, restaurantId uint64) (int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return http.StatusNotFound, nil
	}

	err := ru.RestaurantPostgresRepository.DeleteFavoriteRestaurant(ctx, curUser.Id, restaurantId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (ru *restaurantUsecase) RestaurantSearchPost(ctx context.Context, searchPattern string) (models.RestaurantList, int, error) {
	restaurant, err := ru.RestaurantPostgresRepository.SearchRestaurant(ctx, searchPattern)
	if err != nil {
		return models.RestaurantList{}, http.StatusInternalServerError, err
	}
	return restaurant, http.StatusOK, nil
}
