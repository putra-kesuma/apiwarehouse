package controllers

import (
	"apiwarehouse/middleware"
	"apiwarehouse/models"
	"apiwarehouse/usecases"
	"apiwarehouse/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ItemHandler struct {
	ItemUseCase usecases.ItemUseCase
}

func (h ItemHandler) ListItem(w http.ResponseWriter, r *http.Request) {
	item, err := h.ItemUseCase.GetItem()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem, err := json.Marshal(
		utils.Response(http.StatusOK,"Showing Data successfully", item))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfItem)
}

func (h ItemHandler) InsertItem(w http.ResponseWriter, r *http.Request) {
	item := new(models.Item)
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.Write([]byte("can't decode"))
	}else {
		errUsecase := h.ItemUseCase.InsertItem(item)
		if errUsecase != nil {
			//byteOfItem, _ := json.Marshal(utils.ErrorResponse(http.StatusNoContent,errUsecase))
			//w.Header().Set("Content-Type", "application/json")
			//w.Write(byteOfItem)
			fmt.Println(errUsecase)
			w.Write([]byte(fmt.Sprintf("%v",errUsecase)))
		} else {
			byteOfItem, _ := json.Marshal(utils.OtherResponse(http.StatusOK,"Insert Success"))
			w.Header().Set("Content-Type", "application/json")
			w.Write(byteOfItem)
		}
	}
}

func (h ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	err := h.ItemUseCase.UpdateItem(r)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem,err := json.Marshal(utils.OtherResponse(http.StatusOK,"Update Successfuly"))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfItem)
}

func (h ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := h.ItemUseCase.DeleteItem(&id)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem,err := json.Marshal(utils.OtherResponse(http.StatusOK,"Delete Successfuly"))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfItem)
}

func ItemController(r *mux.Router, model usecases.ItemUseCase){
	ItemHandler := ItemHandler{model}
	sub := r.PathPrefix("").Subrouter()
	sub.Use(middleware.AuthMiddleware)
	sub.HandleFunc("/item", ItemHandler.ListItem).Methods(http.MethodGet)
	sub.HandleFunc("/item", ItemHandler.InsertItem).Methods(http.MethodPost)
	sub.HandleFunc("/item", ItemHandler.UpdateItem).Methods(http.MethodPut)
	sub.HandleFunc("/item/{id}", ItemHandler.DeleteItem).Methods(http.MethodDelete)
}