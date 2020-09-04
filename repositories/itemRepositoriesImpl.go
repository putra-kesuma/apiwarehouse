package repositories

import (
	"apiwarehouse/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

//struct for connection db
type ItemRepoImp struct {
	db *sql.DB
}

func (i ItemRepoImp) GetItemById(id *int) (*models.Item, error) {
	//make var for contain struct warehouse
	dataItem := models.Item{}
	//prepare query
	query := fmt.Sprint("SELECT * FROM view_item WHERE id_item =",*id)
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
		dataItem =  item
	}

	return &dataItem, nil
}

func (i ItemRepoImp) GetCountItem() *float64 {
	rows, err := i.db.Query("SELECT COUNT(*) FROM view_item")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var countRow float64

	for rows.Next() {
		if err := rows.Scan(&countRow); err != nil {
			log.Fatal(err)
		}
	}

	return &countRow
}

func (i ItemRepoImp) GetAllItem() ([]*models.Item, error,*float64) {
	var page,mulai,halPerPage,halaman float64
	halaman=1
	halPerPage = 5
	resultCount:= math.Ceil(*i.GetCountItem()/halPerPage)
	if page == 0 {
		page = 1
	} else {
		page = halaman
	}

	if page>1 {
		mulai = (page * halPerPage) - halPerPage
	} else {
		mulai = 0
	}
	//fmt.Println("jumlah num row",*i.GetCountItem())
	//make var for contain struct item
	dataItem := []*models.Item{}
	//prepare query
	query := fmt.Sprint("SELECT * FROM view_item LIMIT ",mulai,",",halPerPage)
	fmt.Println(resultCount)
	data, err := i.db.Query(query)

	//check error when exec the query
	if err != nil {
		return nil, err,nil
	}
	//for get all data, use looping

	for data.Next() {
		//prepare contain for row data assign to item struct
		item := models.Item{}
		//scan data
		var err = data.Scan(&item.IdItem, &item.Name,
			&item.Dimension, &item.IdTypeItem,&item.NameItem)
		if err != nil {
			return nil, err,nil
		}
		dataItem = append(dataItem, &item)
	}

	return dataItem, nil,&resultCount
}


func (i ItemRepoImp) InsertItem(item *models.Item) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}

	query := `insert into m_item (name,dimensions,id_typeitem)
								value (?,?,?);`

	stmt, err := i.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Print(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(item.Name,item.Dimension,item.IdTypeItem); err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	return tx.Commit()
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