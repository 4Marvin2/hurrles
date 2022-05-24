package repository

import (
	"context"
	"fmt"
	"hurrles/config"
	"hurrles/internal/models"
	"log"

	"github.com/jackc/pgx"
)

type IRestaurantRepository interface {
	GetRestaurants(context.Context) (models.RestaurantList, error)
	GetRestaurantPlacesByFloor(context.Context, models.PlaceParameters) (models.PlaceList, error)
	GetRestaurantAllPlaces(context.Context, uint64) (models.PlaceList, error)
	GetRestaurantBookedPlaces(context.Context, models.PlaceParameters) (models.PlaceList, error)
	GetRestaurantMenuById(context.Context, uint64) (models.DishList, error)
	CreateFavoriteRestaurant(context.Context, uint64, uint64) error
	GetRestaurantFavorite(context.Context, uint64) (models.RestaurantList, error)
	DeleteFavoriteRestaurant(context.Context, uint64, uint64) error
	SearchRestaurant(context.Context, string) (models.RestaurantList, error)
}

type restaurantRepository struct {
	Conn *pgx.ConnPool
}

func NewPostgresRestaurantRepository(config config.PostgresConfig) (IRestaurantRepository, error) {
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

	return &restaurantRepository{pool}, nil
}

func (rr *restaurantRepository) GetRestaurants(ctx context.Context) (models.RestaurantList, error) {
	rows, err := rr.Conn.Query(
		"SELECT id, title, description, address, metro, number, open_time, close_time, kitchen, img, floors FROM restaurants;",
	)
	if err != nil {
		return models.RestaurantList{}, err
	}
	defer rows.Close()

	var restaurants models.RestaurantList
	var curRestaurant models.Restaurant
	for rows.Next() {
		err := rows.Scan(
			&curRestaurant.Id,
			&curRestaurant.Title,
			&curRestaurant.Description,
			&curRestaurant.Address,
			&curRestaurant.Metro,
			&curRestaurant.Number,
			&curRestaurant.OpenTime,
			&curRestaurant.CloseTime,
			&curRestaurant.Kitchen,
			&curRestaurant.Img,
			&curRestaurant.Floors,
		)
		if err != nil {
			return models.RestaurantList{}, err
		}
		restaurants = append(restaurants, curRestaurant)
	}
	if err := rows.Err(); err != nil {
		return models.RestaurantList{}, err
	}
	return restaurants, nil
}

func (rr *restaurantRepository) GetRestaurantPlacesByFloor(ctx context.Context, placeParams models.PlaceParameters) (models.PlaceList, error) {
	rows, err := rr.Conn.Query(
		`SELECT id, capacity, number, left_top, right_bottom, width, height, floor
		FROM places
		WHERE restaurant_id = $1 AND floor = $2;`,
		placeParams.RestaurantId,
		placeParams.Floor,
	)
	if err != nil {
		return models.PlaceList{}, err
	}
	defer rows.Close()

	var places models.PlaceList
	var curPlace models.Place
	for rows.Next() {
		err := rows.Scan(
			&curPlace.Id,
			&curPlace.Capacity,
			&curPlace.Number,
			&curPlace.LeftTop,
			&curPlace.RightBottom,
			&curPlace.Width,
			&curPlace.Height,
			&curPlace.Floor,
		)
		if err != nil {
			return models.PlaceList{}, err
		}
		places = append(places, curPlace)
	}
	if err := rows.Err(); err != nil {
		return models.PlaceList{}, err
	}
	return places, nil
}

func (rr *restaurantRepository) GetRestaurantAllPlaces(ctx context.Context, restaurantId uint64) (models.PlaceList, error) {
	rows, err := rr.Conn.Query(
		`SELECT id, capacity, number, left_top, right_bottom, width, height, floor
		FROM places
		WHERE restaurant_id = $1;`,
		restaurantId,
	)
	if err != nil {
		return models.PlaceList{}, err
	}
	defer rows.Close()

	var places models.PlaceList
	var curPlace models.Place
	for rows.Next() {
		err := rows.Scan(
			&curPlace.Id,
			&curPlace.Capacity,
			&curPlace.Number,
			&curPlace.LeftTop,
			&curPlace.RightBottom,
			&curPlace.Width,
			&curPlace.Height,
			&curPlace.Floor,
		)
		if err != nil {
			return models.PlaceList{}, err
		}
		places = append(places, curPlace)
	}
	if err := rows.Err(); err != nil {
		return models.PlaceList{}, err
	}
	return places, nil
}

func (rr *restaurantRepository) GetRestaurantBookedPlaces(ctx context.Context, placeParams models.PlaceParameters) (models.PlaceList, error) {
	rows, err := rr.Conn.Query(
		`SELECT p.id, p.capacity, p.number, p.left_top, p.right_bottom, p.width, p.height, p.floor
		FROM places as p
		LEFT JOIN orders as o
		ON o.place_id = p.id
		WHERE o.start_time <= $1 AND o.end_time > $1
		AND p.restaurant_id = $2;`,
		placeParams.Time,
		placeParams.RestaurantId,
	)

	if err != nil {
		return models.PlaceList{}, err
	}
	defer rows.Close()

	var places models.PlaceList
	var curPlace models.Place
	for rows.Next() {
		err := rows.Scan(
			&curPlace.Id,
			&curPlace.Capacity,
			&curPlace.Number,
			&curPlace.LeftTop,
			&curPlace.RightBottom,
			&curPlace.Width,
			&curPlace.Height,
			&curPlace.Floor,
		)
		if err != nil {
			return models.PlaceList{}, err
		}
		places = append(places, curPlace)
	}
	if err := rows.Err(); err != nil {
		return models.PlaceList{}, err
	}
	return places, nil
}

func (rr *restaurantRepository) GetRestaurantMenuById(ctx context.Context, id uint64) (models.DishList, error) {
	rows, err := rr.Conn.Query(
		"SELECT id, title, description, price FROM dishes WHERE restaurant_id = $1;",
		id,
	)
	if err != nil {
		return models.DishList{}, err
	}
	defer rows.Close()

	var dishes models.DishList
	var curDish models.Dish
	for rows.Next() {
		err := rows.Scan(
			&curDish.Id,
			&curDish.Title,
			&curDish.Description,
			&curDish.Price,
		)
		if err != nil {
			return models.DishList{}, err
		}
		dishes = append(dishes, curDish)
	}
	if err := rows.Err(); err != nil {
		return models.DishList{}, err
	}
	return dishes, nil
}

func (rr *restaurantRepository) CreateFavoriteRestaurant(ctx context.Context, uid uint64, restaurantId uint64) error {
	var id uint64
	err := rr.Conn.QueryRow(
		`INSERT INTO restaurants_users (user_id, restaurant_id)
		VALUES ($1, $2)
		RETURNING id;`,
		uid,
		restaurantId,
	).Scan(&id)

	if err != nil {
		return err
	}
	return nil
}

func (rr *restaurantRepository) GetRestaurantFavorite(ctx context.Context, uid uint64) (models.RestaurantList, error) {
	rows, err := rr.Conn.Query(
		`SELECT r.id, r.title, r.description, r.address, r.metro, r.number, r.open_time, r.close_time, r.kitchen, r.img, r.floors
		FROM restaurants_users AS ru
		JOIN restaurants AS r ON (r.id = ru.restaurant_id)
		WHERE ru.user_id = $1;`,
		uid,
	)
	if err != nil {
		return models.RestaurantList{}, err
	}
	defer rows.Close()

	var restaurants models.RestaurantList
	var curRestaurant models.Restaurant
	for rows.Next() {
		err := rows.Scan(
			&curRestaurant.Id,
			&curRestaurant.Title,
			&curRestaurant.Description,
			&curRestaurant.Address,
			&curRestaurant.Metro,
			&curRestaurant.Number,
			&curRestaurant.OpenTime,
			&curRestaurant.CloseTime,
			&curRestaurant.Kitchen,
			&curRestaurant.Img,
			&curRestaurant.Floors,
		)
		if err != nil {
			return models.RestaurantList{}, err
		}
		restaurants = append(restaurants, curRestaurant)
	}
	if err := rows.Err(); err != nil {
		return models.RestaurantList{}, err
	}
	return restaurants, nil
}

func (rr *restaurantRepository) DeleteFavoriteRestaurant(ctx context.Context, uid uint64, restaurantId uint64) error {
	var id uint64
	err := rr.Conn.QueryRow(
		`DELETE FROM restaurants_users
		WHERE user_id = $1 AND restaurant_id = $2
		RETURNING id;`,
		uid,
		restaurantId,
	).Scan(&id)

	if err != nil {
		return err
	}
	return nil
}

func (rr *restaurantRepository) SearchRestaurant(ctx context.Context, searchPattern string) (models.RestaurantList, error) {
	rows, err := rr.Conn.Query(
		`SELECT id, title, description, address, metro, number, open_time, close_time, kitchen, img
		FROM restaurants
		WHERE LOWER(title) LIKE '%' || $1 || '%';`,
		searchPattern,
	)
	if err != nil {
		return models.RestaurantList{}, err
	}
	defer rows.Close()

	var restaurants models.RestaurantList
	var curRestaurant models.Restaurant
	for rows.Next() {
		err := rows.Scan(
			&curRestaurant.Id,
			&curRestaurant.Title,
			&curRestaurant.Description,
			&curRestaurant.Address,
			&curRestaurant.Metro,
			&curRestaurant.Number,
			&curRestaurant.OpenTime,
			&curRestaurant.CloseTime,
			&curRestaurant.Kitchen,
			&curRestaurant.Img,
		)
		if err != nil {
			return models.RestaurantList{}, err
		}
		restaurants = append(restaurants, curRestaurant)
	}
	if err := rows.Err(); err != nil {
		return models.RestaurantList{}, err
	}
	return restaurants, nil
}
