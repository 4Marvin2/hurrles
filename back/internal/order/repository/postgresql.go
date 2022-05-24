package repository

import (
	"context"
	"fmt"
	"hurrles/config"
	"hurrles/internal/models"
	"log"

	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

type IOrderRepository interface {
	CreateOrder(context.Context, models.Order) (models.Order, error)
	GetOrders(context.Context, uint64) (models.OrderList, error)
	DeleteOrder(context.Context, uint64, uint64) (models.Order, error)
	GetOrderById(context.Context, uint64) (models.Order, error)
	AddDishToOrder(context.Context, uint64, uint64) (int32, error)
	DeleteDishFromOrder(context.Context, uint64, uint64) (int32, error)
}

type orderRepository struct {
	Conn *pgx.ConnPool
}

func NewPostgresOrderRepository(config config.PostgresConfig) (IOrderRepository, error) {
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

	return &orderRepository{pool}, nil
}

func (or *orderRepository) CreateOrder(ctx context.Context, order models.Order) (models.Order, error) {
	var createdOrder models.Order
	err := or.Conn.QueryRow(
		`INSERT INTO orders (user_id, place_id, start_time, end_time, cost)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, place_id, start_time, end_time, cost, created_time, status;`,
		order.UserId,
		order.PlaceId,
		order.StartTime,
		order.EndTime,
		order.Cost,
	).Scan(
		&createdOrder.Id,
		&createdOrder.UserId,
		&createdOrder.PlaceId,
		&createdOrder.StartTime,
		&createdOrder.EndTime,
		&createdOrder.Cost,
		&createdOrder.CreatedTime,
		&createdOrder.Status,
	)

	if err != nil {
		return models.Order{}, err
	}
	return createdOrder, nil
}

func (or *orderRepository) GetOrders(ctx context.Context, uid uint64) (models.OrderList, error) {
	rows, err := or.Conn.Query(
		`SELECT
			o.id,
			o.user_id,
			o.place_id,
			o.start_time,
			o.end_time,
			o.cost,
			o.created_time,
			o.status,
			r.id,
			r.title,
			r.address,
			r.metro,
			p.number,
			p.capacity,
			ARRAY_REMOVE(ARRAY_AGG(d.id), NULL) AS dish_ids,
			ARRAY_REMOVE(ARRAY_AGG(d.title), NULL) AS dish_titles,
			ARRAY_REMOVE(ARRAY_AGG(d.price), NULL) AS dish_prices,
			ARRAY_REMOVE(ARRAY_AGG(dso.number), NULL) AS dish_numbers
		FROM orders AS o
		JOIN places AS p ON (p.id = o.place_id)
		JOIN restaurants AS r ON (r.id = p.restaurant_id)
		LEFT JOIN dishes_orders AS dso ON (dso.order_id = o.id)
		LEFT JOIN dishes AS d ON (d.id = dso.dish_id)
		WHERE user_id = $1 AND o.end_time >= now()
		GROUP BY
			o.id,
			o.user_id,
			o.place_id,
			o.start_time,
			o.end_time,
			o.cost,
			o.created_time,
			r.id,
			r.title,
			r.address,
			r.metro,
			p.number,
			p.capacity
		ORDER BY o.start_time;`,
		uid,
	)
	if err != nil {
		return models.OrderList{}, err
	}
	defer rows.Close()

	var orders models.OrderList
	var curOrder models.Order
	for rows.Next() {
		err := rows.Scan(
			&curOrder.Id,
			&curOrder.UserId,
			&curOrder.PlaceId,
			&curOrder.StartTime,
			&curOrder.EndTime,
			&curOrder.Cost,
			&curOrder.CreatedTime,
			&curOrder.Status,
			&curOrder.RestaurantId,
			&curOrder.RestaurantTitle,
			&curOrder.RestaurantAddress,
			&curOrder.RestaurantMetro,
			&curOrder.PlaceNumber,
			&curOrder.PlaceCapacity,
			pq.Array(&curOrder.DishesIds),
			pq.Array(&curOrder.DishesTitles),
			pq.Array(&curOrder.DishesPrices),
			pq.Array(&curOrder.DishesCounts),
		)
		if err != nil {
			return models.OrderList{}, err
		}
		orders = append(orders, curOrder)
	}
	if err := rows.Err(); err != nil {
		return models.OrderList{}, err
	}
	return orders, nil
}

func (or *orderRepository) DeleteOrder(ctx context.Context, orderId uint64, uid uint64) (models.Order, error) {
	var deletedOrder models.Order
	err := or.Conn.QueryRow(
		`DELETE FROM orders
		WHERE id = $1 AND user_id = $2 
		RETURNING id, user_id, place_id, start_time, end_time, cost, created_time;`,
		orderId,
		uid,
	).Scan(
		&deletedOrder.Id,
		&deletedOrder.UserId,
		&deletedOrder.PlaceId,
		&deletedOrder.StartTime,
		&deletedOrder.EndTime,
		&deletedOrder.Cost,
		&deletedOrder.CreatedTime,
	)

	if err != nil {
		return models.Order{}, err
	}
	return deletedOrder, nil
}

func (or *orderRepository) GetOrderById(ctx context.Context, orderId uint64) (models.Order, error) {
	var createdOrder models.Order
	err := or.Conn.QueryRow(
		`SELECT id, user_id, place_id, start_time, end_time, cost, created_time, status
		FROM orders
		WHERE id = $1;`,
		orderId,
	).Scan(
		&createdOrder.Id,
		&createdOrder.UserId,
		&createdOrder.PlaceId,
		&createdOrder.StartTime,
		&createdOrder.EndTime,
		&createdOrder.Cost,
		&createdOrder.CreatedTime,
		&createdOrder.Status,
	)
	if err != nil {
		return models.Order{}, err
	}
	return createdOrder, nil
}

func (or *orderRepository) AddDishToOrder(ctx context.Context, orderId uint64, dishId uint64) (int32, error) {
	var findedNumberOfDishes int32
	err := or.Conn.QueryRow(
		`SELECT number
		FROM dishes_orders
		WHERE order_id = $1 AND dish_id = $2;`,
		orderId,
		dishId,
	).Scan(
		&findedNumberOfDishes,
	)

	if err != nil && err != pgx.ErrNoRows {
		return -1, err
	} else if err == pgx.ErrNoRows {
		var number int32
		err := or.Conn.QueryRow(
			`INSERT INTO dishes_orders (order_id, dish_id)
			VALUES ($1, $2)
			RETURNING number;`,
			orderId,
			dishId,
		).Scan(
			&number,
		)

		if err != nil {
			return -1, err
		}
		return number, nil
	}

	var number int32
	err = or.Conn.QueryRow(
		`UPDATE dishes_orders
		SET number = number + 1
		WHERE order_id = $1 AND dish_id = $2
		RETURNING number;`,
		orderId,
		dishId,
	).Scan(
		&number,
	)

	if err != nil {
		return -1, err
	}
	return number, nil
}

func (or *orderRepository) DeleteDishFromOrder(ctx context.Context, orderId uint64, dishId uint64) (int32, error) {
	var findedNumberOfDishes int32
	err := or.Conn.QueryRow(
		`SELECT number
		FROM dishes_orders
		WHERE order_id = $1 AND dish_id = $2;`,
		orderId,
		dishId,
	).Scan(
		&findedNumberOfDishes,
	)

	if err != nil {
		return -1, err
	} else if findedNumberOfDishes == 1 {
		var id uint64
		err := or.Conn.QueryRow(
			`DELETE FROM dishes_orders
			WHERE order_id = $1 AND dish_id = $2 
			RETURNING id;`,
			orderId,
			dishId,
		).Scan(
			&id,
		)

		if err != nil {
			return -1, err
		}
		return 0, nil
	}

	var number int32
	err = or.Conn.QueryRow(
		`UPDATE dishes_orders
		SET number = number - 1
		WHERE order_id = $1 AND dish_id = $2
		RETURNING number;`,
		orderId,
		dishId,
	).Scan(
		&number,
	)

	if err != nil {
		return -1, err
	}
	return number, nil

}
