package repositories

import (
	"apiwarehouse/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

//struct for connection db
type WarehouseStorageRepoImp struct {
	db *sql.DB
}

func (w WarehouseStorageRepoImp) GetAllWarehouseStorage() ([]*models.WarehouseStorage, error) {
	//make var for contain struct warehouse
	dataWarehouseStorage := []*models.WarehouseStorage{}
	//prepare query
	query := "SELECT * FROM view_warehousestorage;"
	data, err := w.db.Query(query)
	//check error when exec the query
	if err != nil {
		return nil, err
	}
	//for get all data, use looping
	for data.Next() {
		//prepare contain for row data assign to warehouse struct
		warehousestorage := models.WarehouseStorage{}
		//scan data
		var err = data.Scan(&warehousestorage.IdWarehouseStorage,&warehousestorage.IdWarehouse ,&warehousestorage.NameWarehouse,
			&warehousestorage.Capacity,&warehousestorage.Address, &warehousestorage.NameTypeWarehouse,&warehousestorage.IdItem,
			&warehousestorage.NameItem,&warehousestorage.Dimension,&warehousestorage.NameTypeItem)
		if err != nil {
			return nil, err
		}
		dataWarehouseStorage = append(dataWarehouseStorage, &warehousestorage)
	}
	//return datawarehouse slice of warehouse and error
	return dataWarehouseStorage, nil
}

func (w WarehouseStorageRepoImp) InsertWarehouseStorage(request *http.Request) error {
	dataWarehouseStorage := models.WarehouseStorage{}
	_ = json.NewDecoder(request.Body).Decode(&dataWarehouseStorage) // json ke struct
	tx, _ := w.db.Begin()
	_, err := tx.Exec(`insert into m_warehousestorage(id_warehouse,id_item) value (?,?);`,
		&dataWarehouseStorage.IdWarehouse,&dataWarehouseStorage.IdItem)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (w WarehouseStorageRepoImp) UpdateWarehouseStorage(request *http.Request) error {
	dataWarehouseStorage := models.WarehouseStorage{}
	_ = json.NewDecoder(request.Body).Decode(&dataWarehouseStorage) // json ke struct
	tx, _ := w.db.Begin()
	_, err := tx.Exec(`update m_warehousestorage set id_warehouse=?,id_item=?
						where id_warehousestorage=?;`,
		&dataWarehouseStorage.IdWarehouse,&dataWarehouseStorage.IdItem,&dataWarehouseStorage.IdWarehouseStorage)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (w WarehouseStorageRepoImp) DeleteWarehouseStorage(id *int) error {
	tx, _ := w.db.Begin()
	_, err := tx.Exec(`delete from m_warehousestorage where id_warehousestorage=?`, *id)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func InitWarehouseStorageRepoImpl(db *sql.DB) WarehouseStorageRepository  {
	return &WarehouseStorageRepoImp{db}
}