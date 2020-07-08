package usecases

import (
	"apiwarehouse/models"
	"net/http"
)

type UserUsecase interface {

	GetAllUser() ([]*models.User,error)
	InsertUser(*models.User) error
	UpdateUser(*http.Request) error
	DeleteUser(*int) error
}
