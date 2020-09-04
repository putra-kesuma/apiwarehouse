package repositories

import (
	"apiwarehouse/models"
)

type CategoryItemRepository interface {
	//blueprint for item
	GetAllCategoryItem() ([]*models.Item,error)
}
