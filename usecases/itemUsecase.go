package usecases

import (
	"apiwarehouse/models"
	"net/http"
)

type ItemUseCase interface {
	GetItem()  ([]*models.Item, error,*float64)
	GetItemById(*int)  (*models.Item, error)
	InsertItem(*models.Item) error
	UpdateItem(*http.Request) error
	DeleteItem(*int) error
}