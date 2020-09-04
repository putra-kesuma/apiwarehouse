package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
	"errors"
	"net/http"
)

type ItemUsecaseImpl struct {
	itemRepo repositories.ItemRepository
}

func (i ItemUsecaseImpl) GetItemById(id *int) (*models.Item, error) {
	item, err:= i.itemRepo.GetItemById(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i ItemUsecaseImpl) InsertItem(item *models.Item) error {
	if item.IdTypeItem == 0 || item.Dimension == 0 || item.Name=="" {
		return errors.New("can't null please check")
	} else {
		err := i.itemRepo.InsertItem(item)
		if err != nil {
			return err
		}
		return nil
	}
}

func (i ItemUsecaseImpl) GetItem() ([]*models.Item, error,*float64) {
	item, err, countRow := i.itemRepo.GetAllItem()
	if err != nil {
		return nil, err,nil
	}

	return item, nil, countRow
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