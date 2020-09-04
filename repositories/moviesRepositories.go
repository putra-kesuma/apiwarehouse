package repositories

import "apiwarehouse/models"

type MoviesRepository interface {
	//blueprint for item
	GetAllMovies() ([]*models.Movies,error,*float64)
	GetMoviesById(*int)  (*models.Movies,error)
	InsertMovies(*models.Movies) error
	GetCountMovies() (*float64)
}