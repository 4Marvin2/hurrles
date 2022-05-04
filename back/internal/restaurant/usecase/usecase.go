package usecase

import (
	"context"
	"hurrles/internal/models"
	"hurrles/internal/restaurant/repository"
	"net/http"
	"time"
)

type IRestaurantUsecase interface {
	RestaurantsGet(context.Context) (models.RestaurantList, int, error)
	RestaurantIdPlacesGet(context.Context, models.PlaceParameters) (models.PlaceList, int, error)
	RestaurantIdMenuGet(context.Context, uint64) (models.DishList, int, error)
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

func (ru *restaurantUsecase) RestaurantIdMenuGet(ctx context.Context, id uint64) (models.DishList, int, error) {
	dishes, err := ru.RestaurantPostgresRepository.GetRestaurantMenuById(ctx, id)
	if err != nil {
		return models.DishList{}, http.StatusInternalServerError, err
	}
	return dishes, http.StatusOK, nil
}
