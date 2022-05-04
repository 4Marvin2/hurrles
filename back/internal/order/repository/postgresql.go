package repository

import (
	"context"
	"fmt"
	"hurrles/config"
	"hurrles/internal/models"
	"log"

	"github.com/jackc/pgx"
)

type IOrderRepository interface {
	CreateOrder(context.Context, models.Order) (models.Order, error)
	GetOrders(context.Context, uint64) (models.OrderList, error)
	DeleteOrder(context.Context, uint64, uint64) (models.Order, error)
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
		RETURNING id, user_id, place_id, start_time, end_time, cost, created_time;`,
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
	)

	if err != nil {
		return models.Order{}, err
	}
	return createdOrder, nil
}

func (or *orderRepository) GetOrders(ctx context.Context, uid uint64) (models.OrderList, error) {
	rows, err := or.Conn.Query(
		`SELECT id, user_id, place_id, start_time, end_time, cost, created_time
		FROM orders
		WHERE user_id = $1 AND end_time >= now();`,
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
		)
		if err != nil {
			return models.OrderList{}, err
		}
		orders = append(orders, curOrder)
	}
	if err := rows.Err(); err != nil {
		return models.OrderList{}, err
	}
	fmt.Println(orders)
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