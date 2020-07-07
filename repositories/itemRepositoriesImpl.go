package repositories

import (
	"apiwarehouse/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

//struct for connection db
type ItemRepoImp struct {
	db *sql.DB
}

func (i ItemRepoImp) GetAllItem() ([]*models.Item, error) {
	//make var for contain struct item
	dataItem := []*models.Item{}
	//prepare query
	query := "SELECT * FROM view_item;"
	data, err := i.db.Query(query)
	//check error when exec the query
	if err != nil {
		return nil, err
	}
	//for get all data, use looping
	for data.Next() {
		//prepare contain for row data assign to item struct
		item := models.Item{}
		//scan data
		var err = data.Scan(&item.IdItem, &item.Name,
			&item.Dimension, &item.IdTypeItem,&item.NameItem)
		if err != nil {
			return nil, err
		}
		dataItem = append(dataItem, &item)
	}
	return dataItem, nil
}

func (i ItemRepoImp) InsertItem(request *http.Request) error {
	dataItem := models.Item{}
	_ = json.NewDecoder(request.Body).Decode(&dataItem) // json ke struct
	tx, _ := i.db.Begin()
	_, err := tx.Exec(`insert into m_item (name,dimensions,id_typeitem)
								value (?,?,?);`,
		&dataItem.Name, &dataItem.Dimension,
		&dataItem.IdTypeItem)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (i ItemRepoImp) UpdateItem(request *http.Request) error {
	dataItem := models.Item{}
	_ = json.NewDecoder(request.Body).Decode(&dataItem) // json ke struct
	tx, _ := i.db.Begin()
	_, err := tx.Exec(`update m_item set name=?,dimensions=?,
						id_typeitem=?
						where id_item=?;`,
		&dataItem.Name, &dataItem.Dimension,
		&dataItem.IdTypeItem, dataItem.IdItem)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (i ItemRepoImp) DeleteItem(id *int) error {
	tx, _ := i.db.Begin()
	_, err := tx.Exec(`delete from m_item where id_item=?`, *id)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func InitItemRepoImpl(db *sql.DB) ItemRepository  {
	return &ItemRepoImp{db}
}