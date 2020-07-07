package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
	"net/http"
)

type WarehouseUsecaseImpl struct {
	warehouseRepo repositories.WarehouseRepository
}

func (w WarehouseUsecaseImpl) GetWarehouse() ([]*models.Warehouse, error) {
	warehouse, err := w.warehouseRepo.GetAllWarehouse()
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (w WarehouseUsecaseImpl) InsertWarehouse(request *http.Request) error {
	err := w.warehouseRepo.InsertWarehouse(request)
	if err != nil {
		return err
	}
	return nil
}

func (w WarehouseUsecaseImpl) UpdateWarehouse(request *http.Request) error {
	err := w.warehouseRepo.UpdateWarehouse(request)
	if err != nil {
		return err
	}
	return nil
}

func (w WarehouseUsecaseImpl) DeleteWarehouse(id *int) error {
	err := w.warehouseRepo.DeleteWarehouse(id)
	if err != nil {
		return err
	}
	return nil
}

func InitWarehouseUsecase(warehuoseRepo repositories.WarehouseRepository) WarehouseUseCase{
	return &WarehouseUsecaseImpl{warehuoseRepo}
}