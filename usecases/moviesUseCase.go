package usecases

import (
	"apiwarehouse/models"
)

type MoviesUseCase interface {
	GetMovies()  ([]*models.Movies, error,*float64)
	GetMoviesById(*int)  (*models.Movies, error)
	InsertMovies(*models.Movies) error
}