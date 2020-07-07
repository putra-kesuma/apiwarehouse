package repositories

import (
	"apiwarehouse/models"
	"database/sql"
)

type ReportWarehouseRepoImp struct {
	db *sql.DB
}

func (r ReportWarehouseRepoImp) GetAllReportWarehouse() ([]*models.ReportWarehouse, error) {
	//make var for contain struct warehouse
	dataReportWarehouse := []*models.ReportWarehouse{}
	//prepare query
	query := "SELECT * FROM view_reportwarehouse;"
	data, err := r.db.Query(query)
	//check error when exec the query
	if err != nil {
		return nil, err
	}
	//for get all data, use looping
	for data.Next() {
		//prepare contain for row data assign to warehouse struct
		reportWarehouse := models.ReportWarehouse{}
		//scan data
		var err = data.Scan(&reportWarehouse.NameWarehouse,&reportWarehouse.TotalItem)
		if err != nil {
			return nil, err
		}
		dataReportWarehouse = append(dataReportWarehouse, &reportWarehouse)
	}
	//return datawarehouse slice of warehouse and error
	return dataReportWarehouse, nil
}

func InitReportWarehouseRepoImpl(db *sql.DB) ReportWarehouseRepository  {
	return &ReportWarehouseRepoImp{db}
}