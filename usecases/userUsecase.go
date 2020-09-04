package usecases

import (
	"apiwarehouse/models"
)

type UserUsecase interface {
	GetUser() ([]*models.User,error)
	InsertUser(*models.User) error
	LoginUser(*models.User) error
	UpdateUser(*models.User) error
	DeleteUser(*int) error
}
