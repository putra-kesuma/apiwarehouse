package repositories

import (
	"apiwarehouse/models"
)

type UserRepository interface {
	//blueprint for user
	GetAllUser() ([]*models.User,error)
	InsertUser(*models.User) error
	LoginUser(*models.User) error
	UpdateUser(*models.User) error
	DeleteUser(*int) error
}
