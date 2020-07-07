package usecases

import (
	"apiwarehouse/models"
	"net/http"
)

type ItemUseCase interface {
	GetItem()  ([]*models.Item, error)
	InsertItem(*http.Request) error
	UpdateItem(*http.Request) error
	DeleteItem(*int) error
}