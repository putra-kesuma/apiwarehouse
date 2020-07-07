package usecases

import (
	"apiwarehouse/models"
	"net/http"
)

type WarehouseUseCase interface {
	GetWarehouse()  ([]*models.Warehouse, error)
	InsertWarehouse(*http.Request) error
	UpdateWarehouse(*http.Request) error
	DeleteWarehouse(*int) error
}

