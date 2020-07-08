package usecases

import (
	"apiwarehouse/models"
	"net/http"
)

type WarehouseStorageUseCase interface {
	GetWarehouseStorage()  ([]*models.WarehouseStorage, error)
	InsertWarehouseStorage(*models.WarehouseStorage) error
	UpdateWarehouseStorage(*http.Request) error
	DeleteWarehouseStorage(*int) error
}
