package usecase

import (
	"context"
	"fmt"
	"hurrles/config"
	"hurrles/internal/models"
	"hurrles/internal/pkg/hasher"
	"hurrles/internal/user/repository"
	"net/http"
	"time"

	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	GetUserById(context.Context, uint64) (models.User, int, error)
	LoginUser(context.Context, models.UserCredentials) (models.User, int, error)
	LoginAdminUser(context.Context, models.UserCredentials) (models.User, int, error)
	LoginRestaurantUser(context.Context, models.UserCredentials) (models.User, int, error)
	SignupUser(context.Context, models.User) (models.User, int, error)
	GetUserFromCtx(context.Context) (models.User, int, error)
	EditUser(context.Context, models.User) (models.User, int, error)
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

func (uu *userUsecase) LoginUser(ctx context.Context, credentials models.UserCredentials) (models.User, int, error) {
	user, err := uu.UserPostgresRepository.GetUserByEmail(ctx, credentials.Email)
	if err == pgx.ErrNoRows {
		return models.User{}, http.StatusNotFound, fmt.Errorf("user with email %s not found", credentials.Email)
	} else if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	isVerify, err := hasher.ComparePasswords(user.Password, credentials.Password)
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		return models.User{}, http.StatusInternalServerError, err
	}

	if !isVerify {
		return models.User{}, http.StatusForbidden, fmt.Errorf("wrong password for user %s", credentials.Email)
	}

	return user, http.StatusOK, nil
}

func (uu *userUsecase) LoginAdminUser(ctx context.Context, credentials models.UserCredentials) (models.User, int, error) {
	user, err := uu.UserPostgresRepository.GetUserByEmail(ctx, credentials.Email)
	if err == pgx.ErrNoRows {
		return models.User{}, http.StatusNotFound, fmt.Errorf("user with email %s not found", credentials.Email)
	} else if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	isVerify, err := hasher.ComparePasswords(user.Password, credentials.Password)
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		return models.User{}, http.StatusInternalServerError, err
	}

	if !isVerify || !user.IsAdmin {
		return models.User{}, http.StatusForbidden, fmt.Errorf("wrong password for user %s", credentials.Email)
	}

	return user, http.StatusOK, nil
}

func (uu *userUsecase) LoginRestaurantUser(ctx context.Context, credentials models.UserCredentials) (models.User, int, error) {
	user, err := uu.UserPostgresRepository.GetUserByEmail(ctx, credentials.Email)
	if err == pgx.ErrNoRows {
		return models.User{}, http.StatusNotFound, fmt.Errorf("user with email %s not found", credentials.Email)
	} else if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	isVerify, err := hasher.ComparePasswords(user.Password, credentials.Password)
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		return models.User{}, http.StatusInternalServerError, err
	}

	if !isVerify || !user.IsRestaurant {
		return models.User{}, http.StatusForbidden, fmt.Errorf("wrong password for user %s", credentials.Email)
	}

	return user, http.StatusOK, nil
}

func (uu *userUsecase) SignupUser(ctx context.Context, user models.User) (models.User, int, error) {
	_, err := uu.UserPostgresRepository.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return models.User{}, http.StatusConflict, err
	} else if err != nil && err != pgx.ErrNoRows {
		return models.User{}, http.StatusInternalServerError, err
	}

	hashedPswd, err := hasher.HashAndSalt(user.Password)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}
	user.Password = hashedPswd

	createdUser, err := uu.UserPostgresRepository.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	return createdUser, http.StatusOK, nil
}

func (uu *userUsecase) GetUserById(ctx context.Context, uid uint64) (models.User, int, error) {
	user, err := uu.UserPostgresRepository.GetUserById(ctx, uid)
	if err == pgx.ErrNoRows {
		return models.User{}, http.StatusNotFound, fmt.Errorf("user with id %d not found", uid)
	} else if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func (uu *userUsecase) GetUserFromCtx(ctx context.Context) (models.User, int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return models.User{}, http.StatusNotFound, nil
	}

	return curUser, http.StatusOK, nil
}

func (uu *userUsecase) EditUser(ctx context.Context, user models.User) (models.User, int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return models.User{}, http.StatusNotFound, nil
	}

	if curUser.Email != user.Email {
		return models.User{}, http.StatusNotFound, nil
	}
	user.Id = curUser.Id

	updatedUser, err := uu.UserPostgresRepository.UpdateUser(ctx, user)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	return updatedUser, http.StatusOK, nil
}
