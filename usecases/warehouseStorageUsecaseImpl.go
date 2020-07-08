package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
	"errors"
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


func (w WarehouseStorageUsecaseImpl) InsertWarehouseStorage(ws *models.WarehouseStorage) error {
	if ws.IdWarehouse == 0 || ws.IdItem == 0 {
		return errors.New("can't null please check")
	} else {
		err := w.warehouseStorageRepo.InsertWarehouseStorage(ws)
		if err != nil {
			return err
		}
		return nil
	}
}

//func (w WarehouseStorageUsecaseImpl) InsertWarehouseStorage(request *http.Request) error {
//	err := w.warehouseStorageRepo.InsertWarehouseStorage(request)
//	if err != nil {
//		return err
//	}
//
//	dataWarehouseStorage := models.WarehouseStorage{}
//	_ = json.NewDecoder(request.Body).Decode(&dataWarehouseStorage) // json ke struct
//		errValid := utils.ValdNotNull(&dataWarehouseStorage)
//		if errValid != nil {
//			return errValid
//		}
//
//		return nil
//
//}

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

