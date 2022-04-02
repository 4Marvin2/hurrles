package user

import "hurrles/internal/user/models"

type IUserUsecase interface {
	CreateUser(models.User) (models.User, error)
}
