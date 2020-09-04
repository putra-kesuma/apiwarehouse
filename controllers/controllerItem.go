package controllers

import (
	"apiwarehouse/models"
	"apiwarehouse/usecases"
	"apiwarehouse/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ItemHandler struct {
	ItemUseCase usecases.ItemUseCase
}

func (h ItemHandler) ListItem(w http.ResponseWriter, r *http.Request) {

	item, err,countRow := h.ItemUseCase.GetItem()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem, err := json.Marshal(
		utils.ListPagesResponse(http.StatusOK, "Showing Data successfully", item,*countRow))
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
	} else {
		errUsecase := h.ItemUseCase.InsertItem(item)
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

func (h ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	err := h.ItemUseCase.UpdateItem(r)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem, err := json.Marshal(utils.OtherResponse(http.StatusOK, "Update Successfuly"))
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
	byteOfItem, err := json.Marshal(utils.OtherResponse(http.StatusOK, "Delete Successfuly"))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfItem)
}

func (h ItemHandler) ListItemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	item, err := h.ItemUseCase.GetItemById(&id)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfItem, err := json.Marshal(item)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfItem)
}

func ItemController(r *mux.Router, model usecases.ItemUseCase) {
	ItemHandler := ItemHandler{model}
	sub := r.PathPrefix("").Subrouter()
	// sub.Use(middleware.AuthMiddleware)
	sub.HandleFunc("/item", ItemHandler.ListItem).Methods(http.MethodGet)
	sub.HandleFunc("/item/{id}", ItemHandler.ListItemById).Methods(http.MethodGet)
	sub.HandleFunc("/item", ItemHandler.InsertItem).Methods(http.MethodPost)
	sub.HandleFunc("/item", ItemHandler.UpdateItem).Methods(http.MethodPut)
	sub.HandleFunc("/item/{id}", ItemHandler.DeleteItem).Methods(http.MethodDelete)
}
