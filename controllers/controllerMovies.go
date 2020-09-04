package controllers

import (
	"apiwarehouse/models"
	"apiwarehouse/usecases"
	"apiwarehouse/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MoviesHandler struct {
	MoviesUseCase usecases.MoviesUseCase
}

func (h MoviesHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies, err,_ := h.MoviesUseCase.GetMovies()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem, err := json.Marshal( movies)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfItem)
}

func (h MoviesHandler) ListMoviesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	movies, err := h.MoviesUseCase.GetMoviesById(&id)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem, err := json.Marshal(movies)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfItem)
}

func (h MoviesHandler) InsertMovies(w http.ResponseWriter, r *http.Request) {
	movies := new(models.Movies)
	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		w.Write([]byte("can't decode"))
	} else {
		errUsecase := h.MoviesUseCase.InsertMovies(movies)
		if errUsecase != nil {

			fmt.Println(errUsecase)
			w.Write([]byte(fmt.Sprintf("%v", errUsecase)))
		} else {
			byteOfItem, _ := json.Marshal(utils.OtherResponse(http.StatusOK, "Insert Success"))
			w.Header().Set("Content-Type", "application/json")
			w.Write(byteOfItem)
		}
	}
}

func MoviesController(r *mux.Router, model usecases.MoviesUseCase) {
	MoviesHandler := MoviesHandler{model}
	sub := r.PathPrefix("").Subrouter()
	// sub.Use(middleware.AuthMiddleware)
	sub.HandleFunc("/movies", MoviesHandler.ListMovies).Methods(http.MethodGet)
	sub.HandleFunc("/movies/{id}", MoviesHandler.ListMoviesById).Methods(http.MethodGet)
	sub.HandleFunc("/movies", MoviesHandler.InsertMovies).Methods(http.MethodPost)
}
