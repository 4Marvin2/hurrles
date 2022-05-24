package usecase

import (
	"context"
	"errors"
	"hurrles/internal/admin/repository"
	"hurrles/internal/models"
	"net/http"
	"time"

	"github.com/jackc/pgx"
)

type IAdminUsecase interface {
	RestaurantCreatePost(context.Context, models.Restaurant) (models.Restaurant, int, error)
	RestaurantUpdatePut(context.Context, models.Restaurant) (models.Restaurant, int, error)

	DishCreatePost(context.Context, models.Dish) (models.Dish, int, error)
	DishUpdatePut(context.Context, models.Dish) (models.Dish, int, error)

	PlaceCreatePost(context.Context, models.Place) (models.Place, int, error)
	PlaceUpdatePut(context.Context, models.Place) (models.Place, int, error)
	AllPlacesUpdatePut(ctx context.Context, places []models.Place) (models.PlaceList, int, error)
	PlaceDropDelete(context.Context, uint64) (models.Place, int, error)
}

type adminUsecase struct {
	AdminPostgresRepository repository.IAdminRepository
	Timeout                 time.Duration
}

var (
	ErrPlaceAlreadyExists = errors.New("place already exists")
)

func NewAdminUsecase(ar repository.IAdminRepository, timeout time.Duration) IAdminUsecase {
	return &adminUsecase{
		AdminPostgresRepository: ar,
		Timeout:                 timeout,
	}
}

//
// RESTAURANTS
//

func (au *adminUsecase) RestaurantCreatePost(ctx context.Context, restaurant models.Restaurant) (models.Restaurant, int, error) {
	_, err := au.AdminPostgresRepository.GetRestaurantByTitleAndAddress(ctx, restaurant.Title, restaurant.Address)
	if err == nil {
		return models.Restaurant{}, http.StatusConflict, err
	} else if err != nil && err != pgx.ErrNoRows {
		return models.Restaurant{}, http.StatusInternalServerError, err
	}

	createdRestaurant, err := au.AdminPostgresRepository.CreateRestaurant(ctx, restaurant)
	if err != nil {
		return models.Restaurant{}, http.StatusInternalServerError, err
	}
	return createdRestaurant, http.StatusOK, nil
}

func (au *adminUsecase) RestaurantUpdatePut(ctx context.Context, restaurant models.Restaurant) (models.Restaurant, int, error) {
	_, err := au.AdminPostgresRepository.GetRestaurantById(ctx, restaurant.Id)
	if err != nil {
		return models.Restaurant{}, http.StatusNotFound, err
	}

	updatedRestaurant, err := au.AdminPostgresRepository.UpdateRestaurant(ctx, restaurant)
	if err != nil {
		return models.Restaurant{}, http.StatusInternalServerError, err
	}
	return updatedRestaurant, http.StatusOK, nil
}

//
// DISHES
//

func (au *adminUsecase) DishCreatePost(ctx context.Context, dish models.Dish) (models.Dish, int, error) {
	_, err := au.AdminPostgresRepository.GetDishByTitleAndRestaurantId(ctx, dish.Title, dish.RestaurantId)
	if err == nil {
		return models.Dish{}, http.StatusConflict, err
	} else if err != nil && err != pgx.ErrNoRows {
		return models.Dish{}, http.StatusInternalServerError, err
	}

	createdDish, err := au.AdminPostgresRepository.CreateDish(ctx, dish)
	if err != nil {
		return models.Dish{}, http.StatusInternalServerError, err
	}
	return createdDish, http.StatusOK, nil
}

func (au *adminUsecase) DishUpdatePut(ctx context.Context, dish models.Dish) (models.Dish, int, error) {
	_, err := au.AdminPostgresRepository.GetDishById(ctx, dish.Id)
	if err != nil {
		return models.Dish{}, http.StatusNotFound, err
	}

	updatedDish, err := au.AdminPostgresRepository.UpdateDish(ctx, dish)
	if err != nil {
		return models.Dish{}, http.StatusInternalServerError, err
	}
	return updatedDish, http.StatusOK, nil
}

//
// PLACES
//

func (au *adminUsecase) PlaceCreatePost(ctx context.Context, place models.Place) (models.Place, int, error) {
	_, err := au.AdminPostgresRepository.GetPlaceByNumberAndRestaurantIdAndCoordinate(
		ctx,
		place.Number,
		place.RestaurantId,
		place.LeftTop,
		place.RightBottom,
		place.Floor,
	)
	if err == nil {
		return models.Place{}, http.StatusConflict, ErrPlaceAlreadyExists
	} else if err != nil && err != pgx.ErrNoRows {
		return models.Place{}, http.StatusInternalServerError, err
	}

	createdPlace, err := au.AdminPostgresRepository.CreatePlace(ctx, place)
	if err != nil {
		return models.Place{}, http.StatusInternalServerError, err
	}
	return createdPlace, http.StatusOK, nil
}

func (au *adminUsecase) PlaceUpdatePut(ctx context.Context, place models.Place) (models.Place, int, error) {
	_, err := au.AdminPostgresRepository.GetPlaceById(ctx, place.Id)
	if err != nil {
		return models.Place{}, http.StatusNotFound, err
	}

	updatedPlace, err := au.AdminPostgresRepository.UpdatePlace(ctx, place)
	if err != nil {
		return models.Place{}, http.StatusInternalServerError, err
	}
	return updatedPlace, http.StatusOK, nil
}

func (au *adminUsecase) AllPlacesUpdatePut(ctx context.Context, places []models.Place) (models.PlaceList, int, error) {
	updatedPlaces := make([]models.Place, len(places))
	for idx, val := range places {
		_, err := au.AdminPostgresRepository.GetPlaceById(ctx, val.Id)
		if err != nil {
			return nil, http.StatusNotFound, err
		}

		updatedPlace, err := au.AdminPostgresRepository.UpdatePlace(ctx, val)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		updatedPlaces[idx] = updatedPlace
	}

	return updatedPlaces, http.StatusOK, nil
}

func (au *adminUsecase) PlaceDropDelete(ctx context.Context, placeId uint64) (models.Place, int, error) {
	updatedPlace, err := au.AdminPostgresRepository.DeletePlace(ctx, placeId)
	if err != nil {
		return models.Place{}, http.StatusInternalServerError, err
	}
	return updatedPlace, http.StatusOK, nil
}
