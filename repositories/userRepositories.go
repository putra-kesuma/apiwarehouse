package repositories

import (
	"apiwarehouse/models"
	"net/http"
)

type UserRepository interface {
	//blueprint for item
	GetAllUser() ([]*models.User,error)
	InsertUser(*models.User) error
	UpdateUser(*http.Request) error
	DeleteUser(*int) error
}
