package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
	"net/http"
)

type WarehouseStorageUsecaseImpl struct {
	warehouseStorageRepo repositories.WarehouseStorageRepository
}

func (w WarehouseStorageUsecaseImpl) GetWarehouseStorage() ([]*models.WarehouseStorage, error) {
	warehouseStorage, err := w.warehouseStorageRepo.GetAllWarehouseStorage()
	if err != nil {
		return nil, err
	}

	return warehouseStorage, nil
}

func (w WarehouseStorageUsecaseImpl) InsertWarehouseStorage(request *http.Request) error {

		err := w.warehouseStorageRepo.InsertWarehouseStorage(request)
		if err != nil {
			return err
		}

		return nil

}

func (w WarehouseStorageUsecaseImpl) UpdateWarehouseStorage(request *http.Request) error {
	err := w.warehouseStorageRepo.UpdateWarehouseStorage(request)
	if err != nil {
		return err
	}
	return nil
}

func (w WarehouseStorageUsecaseImpl) DeleteWarehouseStorage(id *int) error {
	err := w.warehouseStorageRepo.DeleteWarehouseStorage(id)
	if err != nil {
		return err
	}
	return nil
}

func InitWarehouseStorageUsecase(warehuoseStorageRepo repositories.WarehouseStorageRepository) WarehouseStorageUseCase{
	return &WarehouseStorageUsecaseImpl{warehuoseStorageRepo}
}

