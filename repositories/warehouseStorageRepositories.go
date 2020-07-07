package repositories

import (
	"apiwarehouse/models"
	"net/http"
)

type WarehouseStorageRepository interface {
	//blueprint for warehousestorage
	GetAllWarehouseStorage() ([]*models.WarehouseStorage,error)
	InsertWarehouseStorage(*http.Request) error
	UpdateWarehouseStorage(*http.Request) error
	DeleteWarehouseStorage(*int) error
}
