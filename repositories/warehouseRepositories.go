package repositories

import (
	"apiwarehouse/models"
	"net/http"
)

type WarehouseRepository interface {
	//blueprint for warehouse
	GetAllWarehouse() ([]*models.Warehouse,error)
	InsertWarehouse(*http.Request) error
	UpdateWarehouse(*http.Request) error
	DeleteWarehouse(*int) error
}

