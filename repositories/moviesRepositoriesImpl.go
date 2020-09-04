package repositories

import (
	"apiwarehouse/models"
	"database/sql"
	"fmt"
	"log"
	"math"
)

//struct for connection db
type MoviesRepoImp struct {
	db *sql.DB
}

func (m MoviesRepoImp) GetAllMovies() ([]*models.Movies, error, *float64) {
	var page,mulai,halPerPage,halaman float64
	halaman=1
	halPerPage = 5
	resultCount:= math.Ceil(*m.GetCountMovies()/halPerPage)
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
	dataMovies := []*models.Movies{}
	//prepare query
	query := fmt.Sprint("SELECT * FROM m_movies LIMIT ",mulai,",",halPerPage)
	fmt.Println(resultCount)
	data, err := m.db.Query(query)

	//check error when exec the query
	if err != nil {
		return nil, err,nil
	}
	//for get all data, use looping

	for data.Next() {
		//prepare contain for row data assign to item struct
		movies := models.Movies{}
		//scan data
		var err = data.Scan(&movies.IdMovie, &movies.Title,
			&movies.Duration, &movies.ImageUrl,&movies.Synopsis)
		if err != nil {
			return nil, err,nil
		}
		dataMovies = append(dataMovies, &movies)
	}

	return dataMovies, nil,&resultCount
}

func (m MoviesRepoImp) GetMoviesById(id *int) (*models.Movies, error) {
	//make var for contain struct warehouse
	dataMovies := models.Movies{}
	//prepare query
	query := fmt.Sprint("SELECT * FROM m_movies WHERE id =",*id)
	data, err := m.db.Query(query)

	//check error when exec the query
	if err != nil {
		return nil, err
	}
	//for get all data, use looping

	for data.Next() {
		//prepare contain for row data assign to item struct
		movies := models.Movies{}
		//scan data
		var err = data.Scan(&movies.IdMovie, &movies.Title,
			&movies.Duration, &movies.ImageUrl,&movies.Synopsis)
		if err != nil {
			return nil, err
		}
		dataMovies =  movies
	}

	return &dataMovies, nil
}

func (m MoviesRepoImp) InsertMovies(movies *models.Movies) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	query := `insert into m_movies (title,duration,image_url,synopsis)
								value (?,?,?,?);`

	stmt, err := m.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Print(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(movies.Title,movies.Duration,movies.ImageUrl,movies.Synopsis); err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	return tx.Commit()
}

func (m MoviesRepoImp) GetCountMovies() *float64 {
	rows, err := m.db.Query("SELECT COUNT(*) FROM m_movies")
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

func InitMoviesRepoImpl(db *sql.DB) MoviesRepository  {
	return &MoviesRepoImp{db}
}