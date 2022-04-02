package usecase

import (
	"fmt"
	"hurrles/internal/pkg/hasher"
	"hurrles/internal/user/models"
	"hurrles/internal/user/repository"
	"net/http"
	"time"

	"github.com/jackc/pgx"
)

type IUserUsecase interface {
	GetUserById(uid uint64) (models.User, error)
	LoginUser(models.UserCredentials) (models.User, int, error)
	SignupUser(models.User) (models.User, int, error)
}

type userUsecase struct {
	UserPostgresRepository repository.IUserRepository
	Timeout                time.Duration
}

func NewUserUsecase(ur repository.IUserRepository, timeout time.Duration) IUserUsecase {
	return &userUsecase{
		UserPostgresRepository: ur,
		Timeout:                timeout,
	}
}

func (uu *userUsecase) LoginUser(credentials models.UserCredentials) (models.User, int, error) {
	user, err := uu.UserPostgresRepository.GetUserByEmail(credentials.Email)
	if err == pgx.ErrNoRows {
		return models.User{}, http.StatusNotFound, fmt.Errorf("user with email %s not found", credentials.Email)
	} else if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	isVerify, err := hasher.ComparePasswords(user.Password, credentials.Password)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	if !isVerify {
		return models.User{}, http.StatusForbidden, fmt.Errorf("wrong password for user %s", credentials.Email)
	}

	return user, http.StatusOK, nil
}

func (uu *userUsecase) SignupUser(user models.User) (models.User, int, error) {
	user, err := uu.UserPostgresRepository.GetUserByEmail(user.Email)
	if err == nil {
		return models.User{}, http.StatusConflict, err
	} else if err != nil && err != pgx.ErrNoRows {
		return models.User{}, http.StatusInternalServerError, err
	}

	createdUser, err := uu.UserPostgresRepository.CreateUser(user)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	return createdUser, http.StatusOK, nil
}

func (uu *userUsecase) GetUserById(uid uint64) (models.User, error) {
	user, err := uu.UserPostgresRepository.GetUserById(uid)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
