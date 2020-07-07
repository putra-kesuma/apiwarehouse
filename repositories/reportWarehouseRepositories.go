package repositories

import (
	"apiwarehouse/models"
)

type ReportWarehouseRepository interface {
	//blueprint for warehouse
	GetAllReportWarehouse() ([]*models.ReportWarehouse,error)
}
