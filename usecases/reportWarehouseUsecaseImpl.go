package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
)

type ReportWarehouseUsecaseImpl struct {
	reportWarehouseRepo repositories.ReportWarehouseRepository
}

func (r ReportWarehouseUsecaseImpl) GetReportWarehouse() ([]*models.ReportWarehouse, error) {
	reportWarehouse, err := r.reportWarehouseRepo.GetAllReportWarehouse()
	if err != nil {
		return nil, err
	}

	return reportWarehouse, nil
}

func InitReportWarehouseUsecase(reportWarehuoseRepo repositories.ReportWarehouseRepository) ReportWarehouseUseCase{
	return &ReportWarehouseUsecaseImpl{reportWarehuoseRepo}
}