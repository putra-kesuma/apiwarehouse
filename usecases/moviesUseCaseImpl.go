package usecases

import (
	"apiwarehouse/models"
	"apiwarehouse/repositories"
)

type MoviesUsecaseImpl struct {
	moviesRepo repositories.MoviesRepository
}

func (m MoviesUsecaseImpl) GetMovies() ([]*models.Movies, error, *float64) {
	movies, err, countRow := m.moviesRepo.GetAllMovies()
	if err != nil {
		return nil, err,nil
	}

	return movies, nil, countRow
}

func (m MoviesUsecaseImpl) GetMoviesById(id *int) (*models.Movies, error) {
	movies, err:= m.moviesRepo.GetMoviesById(id)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (m MoviesUsecaseImpl) InsertMovies(movies *models.Movies) error {
	err := m.moviesRepo.InsertMovies(movies)
	if err != nil {
		return err
	}
	return nil
}

func InitMoviesUsecase(moviesRepo repositories.MoviesRepository) MoviesUseCase{
	return &MoviesUsecaseImpl{moviesRepo }
}
