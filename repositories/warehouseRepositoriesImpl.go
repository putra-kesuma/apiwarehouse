package repositories

import (
	"apiwarehouse/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

//struct for connection db
type WarehouseRepoImp struct {
	db *sql.DB
}

func (w WarehouseRepoImp) GetAllWarehouse() ([]*models.Warehouse, error) {
	//make var for contain struct warehouse
	dataWarehouse := []*models.Warehouse{}
	//prepare query
	query := "SELECT * FROM view_warehouse;"
	data, err := w.db.Query(query)
	//check error when exec the query
	if err != nil {
		return nil, err
	}
	//for get all data, use looping
	for data.Next() {
		//prepare contain for row data assign to warehouse struct
		warehouse := models.Warehouse{}
		//scan data
		var err = data.Scan(&warehouse.IdWarehouse, &warehouse.Name,
			&warehouse.Capacity,&warehouse.Address, &warehouse.IdTypeWarehouse,&warehouse.NameType)
		if err != nil {
			return nil, err
		}
		dataWarehouse = append(dataWarehouse, &warehouse)
	}
	//return datawarehouse slice of warehouse and error
	return dataWarehouse, nil
}

func (w WarehouseRepoImp) InsertWarehouse(request *http.Request) error {
	dataWarehouse := models.Warehouse{}
	_ = json.NewDecoder(request.Body).Decode(&dataWarehouse) // json ke struct
	tx, _ := w.db.Begin()
	_, err := tx.Exec(`insert into m_warehouse(name,capacity,address,id_typewarehouse)
								value (?,?,?,?);`,
								&dataWarehouse.Name, &dataWarehouse.Capacity,&dataWarehouse.Address,
							&dataWarehouse.IdTypeWarehouse)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (w WarehouseRepoImp) UpdateWarehouse(request *http.Request) error {
	dataWarehouse := models.Warehouse{}
	_ = json.NewDecoder(request.Body).Decode(&dataWarehouse) // json ke struct
	tx, _ := w.db.Begin()
	_, err := tx.Exec(`update m_warehouse set name=?,capacity=?,
						address=?,id_typewarehouse=?
						where id_warehouse=?;`,
						&dataWarehouse.Name, &dataWarehouse.Capacity,&dataWarehouse.Address,
						&dataWarehouse.IdTypeWarehouse, dataWarehouse.IdWarehouse)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (w WarehouseRepoImp) DeleteWarehouse(id *int) error {
	tx, _ := w.db.Begin()
	_, err := tx.Exec(`delete from m_warehouse where id_warehouse=?`, *id)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func InitWarehouseRepoImpl(db *sql.DB) WarehouseRepository  {
	return &WarehouseRepoImp{db}
}