package repository

import (
	"context"
	"fmt"
	"hurrles/config"
	"hurrles/internal/models"
	"log"

	"github.com/jackc/pgx"
)

type IUserRepository interface {
	GetUserById(context.Context, uint64) (models.User, error)
	GetUserByEmail(context.Context, string) (models.User, error)
	CreateUser(context.Context, models.User) (models.User, error)
	UpdateUser(context.Context, models.User) (models.User, error)
}

type userRepository struct {
	Conn *pgx.ConnPool
}

func NewPostgresUserRepository(config config.PostgresConfig) (IUserRepository, error) {
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

	return &userRepository{pool}, nil
}

func (ur *userRepository) GetUserById(ctx context.Context, uid uint64) (models.User, error) {
	var user models.User
	err := ur.Conn.QueryRow(
		"SELECT id, email, password, full_name, number, is_admin, is_restaurant, restaurant FROM users WHERE id = $1;",
		uid,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.Number,
		&user.IsAdmin,
		&user.IsRestaurant,
		&user.Restaurant,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := ur.Conn.QueryRow(
		"SELECT id, email, password, full_name, number, is_admin, is_restaurant, restaurant FROM users WHERE email = $1;",
		email,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.Number,
		&user.IsAdmin,
		&user.IsRestaurant,
		&user.Restaurant,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	var createdUser models.User
	err := ur.Conn.QueryRow(
		`INSERT INTO users (email, password)
		VALUES ($1, $2) RETURNING id, email, password, full_name, number, is_admin, is_restaurant, restaurant;`,
		user.Email,
		user.Password,
	).Scan(
		&createdUser.Id,
		&createdUser.Email,
		&createdUser.Password,
		&createdUser.FullName,
		&createdUser.Number,
		&user.IsAdmin,
		&user.IsRestaurant,
		&user.Restaurant,
	)

	if err != nil {
		return models.User{}, err
	}
	return createdUser, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	var updatedUser models.User
	err := ur.Conn.QueryRow(
		`UPDATE users
		SET (full_name, email, number) = ($2, $3, $4)
		WHERE id = $1
		RETURNING id, email, password, full_name, number, is_admin, is_restaurant, restaurant;`,
		user.Id,
		user.FullName,
		user.Email,
		user.Number,
	).Scan(
		&updatedUser.Id,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.FullName,
		&updatedUser.Number,
		&user.IsAdmin,
		&user.IsRestaurant,
		&user.Restaurant,
	)

	if err != nil {
		return models.User{}, err
	}
	return updatedUser, nil
}
