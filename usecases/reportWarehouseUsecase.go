package usecases

import (
	"apiwarehouse/models"
)

type ReportWarehouseUseCase interface {
	GetReportWarehouse()  ([]*models.ReportWarehouse, error)
}

