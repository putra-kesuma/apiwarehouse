package utils

import (
	"apiwarehouse/models"
	"errors"
)

func ValdNotNull(dataWarehouseStorage *models.WarehouseStorage) error{
	if (*dataWarehouseStorage).IdWarehouse == 0 || (*dataWarehouseStorage).IdItem == 0 {
		return errors.New("Can't empety")
		}else {
		return nil

	}
}
