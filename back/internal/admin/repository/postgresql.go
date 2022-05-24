package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hurrles/config"
	"hurrles/internal/models"
	"log"

	"github.com/jackc/pgx"
)

type IAdminRepository interface {
	GetRestaurantById(context.Context, uint64) (models.Restaurant, error)
	GetRestaurantByTitleAndAddress(context.Context, string, string) (models.Restaurant, error)
	CreateRestaurant(context.Context, models.Restaurant) (models.Restaurant, error)
	UpdateRestaurant(context.Context, models.Restaurant) (models.Restaurant, error)

	GetDishById(context.Context, uint64) (models.Dish, error)
	GetDishByTitleAndRestaurantId(context.Context, string, uint64) (models.Dish, error)
	CreateDish(context.Context, models.Dish) (models.Dish, error)
	UpdateDish(context.Context, models.Dish) (models.Dish, error)

	GetPlaceById(context.Context, uint64) (models.Place, error)
	GetPlaceByNumberAndRestaurantIdAndCoordinate(context.Context, int, uint64, int, int, int) (models.Place, error)
	CreatePlace(context.Context, models.Place) (models.Place, error)
	UpdatePlace(context.Context, models.Place) (models.Place, error)
	DeletePlace(context.Context, uint64) (models.Place, error)
}

type adminRepository struct {
	Conn *pgx.ConnPool
}

func NewPostgresAdminRepository(config config.PostgresConfig) (IAdminRepository, error) {
	ConnStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.User,
		config.DBName,
		config.Password,
		config.Host,
		config.Port)

	pgxConnectionConfig, err := pgx.ParseConnectionString(ConnStr)
	if err != nil {
		log.Fatalf("Invalid config string: %s", err)
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     pgxConnectionConfig,
		MaxConnections: 100,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	})
	if err != nil {
		log.Fatalf("Error %s occurred during connection to database", err)
	}

	return &adminRepository{pool}, nil
}

//
// RESTAURANTS
//

func (ar *adminRepository) GetRestaurantById(ctx context.Context, id uint64) (models.Restaurant, error) {
	var restaurant models.Restaurant
	err := ar.Conn.QueryRow(
		`SELECT id, title, description, address, number, open_time, close_time, kitchen, img, floors
		FROM restaurants
		WHERE id = $1;`,
		id,
	).Scan(
		&restaurant.Id,
		&restaurant.Title,
		&restaurant.Description,
		&restaurant.Address,
		&restaurant.Number,
		&restaurant.OpenTime,
		&restaurant.CloseTime,
		&restaurant.Kitchen,
		&restaurant.Img,
		&restaurant.Floors,
	)
	if err != nil {
		return models.Restaurant{}, err
	}
	return restaurant, nil
}

func (ar *adminRepository) GetRestaurantByTitleAndAddress(ctx context.Context, title string, address string) (models.Restaurant, error) {
	var restaurant models.Restaurant
	err := ar.Conn.QueryRow(
		`SELECT id, title, description, address, number, open_time, close_time, kitchen, img, floors
		FROM restaurants
		WHERE title = $1 AND address = $2;`,
		title,
		address,
	).Scan(
		&restaurant.Id,
		&restaurant.Title,
		&restaurant.Description,
		&restaurant.Address,
		&restaurant.Number,
		&restaurant.OpenTime,
		&restaurant.CloseTime,
		&restaurant.Kitchen,
		&restaurant.Img,
		&restaurant.Floors,
	)
	if err != nil {
		return models.Restaurant{}, err
	}
	return restaurant, nil
}

func (ar *adminRepository) CreateRestaurant(ctx context.Context, restaurant models.Restaurant) (models.Restaurant, error) {
	var createdRestaurant models.Restaurant
	// formattedOpenTime := restaurant.OpenTime.Format("15:04")
	// formattedCloseTime := restaurant.CloseTime.Format("15:04")
	// var openTime, closeTime interface{}
	err := ar.Conn.QueryRow(
		`INSERT INTO restaurants (title, description, address, metro, number, open_time, close_time, kitchen, img, floors)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, title, description, address, metro, number, open_time, close_time, kitchen, img, floors;`,
		restaurant.Title,
		restaurant.Description,
		restaurant.Address,
		restaurant.Metro,
		restaurant.Number,
		restaurant.OpenTime,
		restaurant.CloseTime,
		restaurant.Kitchen,
		restaurant.Img,
		restaurant.Floors,
	).Scan(
		&createdRestaurant.Id,
		&createdRestaurant.Title,
		&createdRestaurant.Description,
		&createdRestaurant.Address,
		&createdRestaurant.Metro,
		&createdRestaurant.Number,
		&createdRestaurant.OpenTime,
		&createdRestaurant.CloseTime,
		&createdRestaurant.Kitchen,
		&createdRestaurant.Img,
		&createdRestaurant.Floors,
	)

	if err != nil {
		return models.Restaurant{}, err
	}
	return createdRestaurant, nil
}

func (ar *adminRepository) UpdateRestaurant(ctx context.Context, restaurant models.Restaurant) (models.Restaurant, error) {
	var updatedRestaurant models.Restaurant
	err := ar.Conn.QueryRow(
		`UPDATE restaurants
		SET (title, description, address, metro, number, open_time, close_time, kitchen, img, floors) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		WHERE id = $1
		RETURNING id, title, description, address, metro, number, open_time, close_time, kitchen, img, floors;`,
		restaurant.Id,
		restaurant.Title,
		restaurant.Description,
		restaurant.Address,
		restaurant.Metro,
		restaurant.Number,
		restaurant.OpenTime,
		restaurant.CloseTime,
		restaurant.Kitchen,
		restaurant.Img,
		restaurant.Floors,
	).Scan(
		&updatedRestaurant.Id,
		&updatedRestaurant.Title,
		&updatedRestaurant.Description,
		&updatedRestaurant.Address,
		&updatedRestaurant.Metro,
		&updatedRestaurant.Number,
		&updatedRestaurant.OpenTime,
		&updatedRestaurant.CloseTime,
		&updatedRestaurant.Kitchen,
		&updatedRestaurant.Img,
		&updatedRestaurant.Floors,
	)

	if err != nil {
		return models.Restaurant{}, err
	}
	return updatedRestaurant, nil
}

//
// DISHES
//

func (ar *adminRepository) GetDishById(ctx context.Context, id uint64) (models.Dish, error) {
	var dish models.Dish
	err := ar.Conn.QueryRow(
		`SELECT id, restaurant_id, title, description, price
		FROM dishes
		WHERE id = $1;`,
		id,
	).Scan(
		&dish.Id,
		&dish.RestaurantId,
		&dish.Title,
		&dish.Description,
		&dish.Price,
	)
	if err != nil {
		return models.Dish{}, err
	}
	return dish, nil
}

func (ar *adminRepository) GetDishByTitleAndRestaurantId(ctx context.Context, title string, restaurantId uint64) (models.Dish, error) {
	var dish models.Dish
	err := ar.Conn.QueryRow(
		`SELECT id, restaurant_id, title, description, price
		FROM dishes
		WHERE title = $1 AND restaurant_id = $2;`,
		title,
		restaurantId,
	).Scan(
		&dish.Id,
		&dish.RestaurantId,
		&dish.Title,
		&dish.Description,
		&dish.Price,
	)
	if err != nil {
		return models.Dish{}, err
	}
	return dish, nil
}

func (ar *adminRepository) CreateDish(ctx context.Context, dish models.Dish) (models.Dish, error) {
	var createdDish models.Dish
	err := ar.Conn.QueryRow(
		`INSERT INTO dishes (restaurant_id, title, description, price)
		VALUES ($1, $2, $3, $4)
		RETURNING id, restaurant_id, title, description, price;`,
		dish.RestaurantId,
		dish.Title,
		dish.Description,
		dish.Price,
	).Scan(
		&createdDish.Id,
		&createdDish.RestaurantId,
		&createdDish.Title,
		&createdDish.Description,
		&createdDish.Price,
	)

	if err != nil {
		return models.Dish{}, err
	}
	return createdDish, nil
}

func (ar *adminRepository) UpdateDish(ctx context.Context, dish models.Dish) (models.Dish, error) {
	var updatedDish models.Dish
	err := ar.Conn.QueryRow(
		`UPDATE dishes
		SET (restaurant_id, title, description, price) = ($2, $3, $4, $5)
		WHERE id = $1
		RETURNING id, restaurant_id, title, description, price;`,
		dish.Id,
		dish.RestaurantId,
		dish.Title,
		dish.Description,
		dish.Price,
	).Scan(
		&updatedDish.Id,
		&updatedDish.RestaurantId,
		&updatedDish.Title,
		&updatedDish.Description,
		&updatedDish.Price,
	)

	if err != nil {
		return models.Dish{}, err
	}
	return updatedDish, nil
}

//
// PLACES
//

func (ar *adminRepository) GetPlaceById(ctx context.Context, id uint64) (models.Place, error) {
	var place models.Place
	err := ar.Conn.QueryRow(
		`SELECT id, restaurant_id, capacity, number, left_top, right_bottom, width, height, floor
		FROM places
		WHERE id = $1;`,
		id,
	).Scan(
		&place.Id,
		&place.RestaurantId,
		&place.Capacity,
		&place.Number,
		&place.LeftTop,
		&place.RightBottom,
		&place.Width,
		&place.Height,
		&place.Floor,
	)
	if err != nil {
		return models.Place{}, err
	}
	return place, nil
}

func (ar *adminRepository) GetPlaceByNumberAndRestaurantIdAndCoordinate(
	ctx context.Context,
	number int,
	restaurantId uint64,
	leftTop int,
	rightBottom int,
	floor int) (models.Place, error) {
	var place models.Place
	err := ar.Conn.QueryRow(
		`SELECT id, restaurant_id, capacity, number, left_top, right_bottom, width, height, floor
		FROM places
		WHERE (number = $1 AND restaurant_id = $2) OR (left_top = $3 AND right_bottom = $4 AND restaurant_id = $2 AND floor=$5);`,
		number,
		restaurantId,
		leftTop,
		rightBottom,
		floor,
	).Scan(
		&place.Id,
		&place.RestaurantId,
		&place.Capacity,
		&place.Number,
		&place.LeftTop,
		&place.RightBottom,
		&place.Width,
		&place.Height,
		&place.Floor,
	)
	if err != nil {
		return models.Place{}, err
	}
	return place, nil
}

func (ar *adminRepository) CreatePlace(ctx context.Context, place models.Place) (models.Place, error) {
	var createdPlace models.Place
	err := ar.Conn.QueryRow(
		`INSERT INTO places (restaurant_id, capacity, number, left_top, right_bottom, width, height, floor)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, restaurant_id, capacity, number, left_top, right_bottom, width, height, floor;`,
		place.RestaurantId,
		place.Capacity,
		place.Number,
		place.LeftTop,
		place.RightBottom,
		place.Width,
		place.Height,
		place.Floor,
	).Scan(
		&createdPlace.Id,
		&createdPlace.RestaurantId,
		&createdPlace.Capacity,
		&createdPlace.Number,
		&createdPlace.LeftTop,
		&createdPlace.RightBottom,
		&createdPlace.Width,
		&createdPlace.Height,
		&createdPlace.Floor,
	)

	if err != nil {
		return models.Place{}, err
	}
	return createdPlace, nil
}

func (ar *adminRepository) UpdatePlace(ctx context.Context, place models.Place) (models.Place, error) {
	var updatedPlace models.Place
	err := ar.Conn.QueryRow(
		`UPDATE places
		SET (capacity, number, left_top, right_bottom, width, height) = ($2, $3, $4, $5, $6, $7)
		WHERE id = $1
		RETURNING id, restaurant_id, capacity, number, left_top, right_bottom, width, height, floor;`,
		place.Id,
		place.Capacity,
		place.Number,
		place.LeftTop,
		place.RightBottom,
		place.Width,
		place.Height,
	).Scan(
		&updatedPlace.Id,
		&updatedPlace.RestaurantId,
		&updatedPlace.Capacity,
		&updatedPlace.Number,
		&updatedPlace.LeftTop,
		&updatedPlace.RightBottom,
		&updatedPlace.Width,
		&updatedPlace.Height,
		&updatedPlace.Floor,
	)

	if err != nil {
		return models.Place{}, err
	}
	return updatedPlace, nil
}

func (ar *adminRepository) DeletePlace(ctx context.Context, placeId uint64) (models.Place, error) {
	var updatedPlace models.Place
	err := ar.Conn.QueryRow(
		`DELETE FROM places WHERE id = $1
		RETURNING id, restaurant_id, capacity, number, left_top, right_bottom, width, height, floor;`,
		placeId,
	).Scan(
		&updatedPlace.Id,
		&updatedPlace.RestaurantId,
		&updatedPlace.Capacity,
		&updatedPlace.Number,
		&updatedPlace.LeftTop,
		&updatedPlace.RightBottom,
		&updatedPlace.Width,
		&updatedPlace.Height,
		&updatedPlace.Floor,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Place{}, nil
		} else {
			return models.Place{}, err
		}
	}
	return updatedPlace, nil
}
