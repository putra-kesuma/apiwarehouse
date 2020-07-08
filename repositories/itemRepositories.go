package repositories

import (
	"apiwarehouse/models"
	"net/http"
)

type ItemRepository interface {
	//blueprint for item
	GetAllItem() ([]*models.Item,error)
	InsertItem(*models.Item) error
	UpdateItem(*http.Request) error
	DeleteItem(*int) error
}
