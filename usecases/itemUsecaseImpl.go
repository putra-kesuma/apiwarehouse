package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
	"net/http"
)

type ItemUsecaseImpl struct {
	itemRepo repositories.ItemRepository
}

func (i ItemUsecaseImpl) GetItem() ([]*models.Item, error) {
	item, err := i.itemRepo.GetAllItem()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i ItemUsecaseImpl) InsertItem(request *http.Request) error {
	err := i.itemRepo.InsertItem(request)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemUsecaseImpl) UpdateItem(request *http.Request) error {
	err := i.itemRepo.UpdateItem(request)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemUsecaseImpl) DeleteItem(id *int) error {
	err :=i.itemRepo.DeleteItem(id)
	if err != nil {
		return err
	}
	return nil
}

func InitItemUsecase(itemRepo repositories.ItemRepository) ItemUseCase{
	return &ItemUsecaseImpl{itemRepo }
}