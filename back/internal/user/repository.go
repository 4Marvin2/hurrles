package user

import "hurrles/internal/user/models"

type IUserRepository interface {
	CreateUser(models.User) (models.User, error)
}
