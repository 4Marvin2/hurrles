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
		"SELECT id, email, password, full_name, number FROM users WHERE id = $1;",
		uid,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.Number,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := ur.Conn.QueryRow(
		"SELECT id, email, password, full_name, number FROM users WHERE email = $1;",
		email,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.Number,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	var createdUser models.User
	err := ur.Conn.QueryRow(
		`INSERT INTO users (email, password, full_name, number)
		VALUES ($1, $2, $3, $4) RETURNING email, password, full_name, number;`,
		user.Email,
		user.Password,
		user.FullName,
		user.Number,
	).Scan(
		&createdUser.Email,
		&createdUser.Password,
		&createdUser.FullName,
		&createdUser.Number,
	)

	if err != nil {
		return models.User{}, err
	}
	return createdUser, nil
}
