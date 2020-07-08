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

type WarehouseStorageHandler struct {
	WarehouseStorageUseCase usecases.WarehouseStorageUseCase
}

func (h WarehouseStorageHandler) ListWarehouseStorage(w http.ResponseWriter, r *http.Request) {
	warehouseStorage, err := h.WarehouseStorageUseCase.GetWarehouseStorage()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}else {
		byteOfWarehouseStorage, err := json.Marshal(
			utils.Response(http.StatusOK,"Showing Data successfully", warehouseStorage))
		if err != nil {
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(byteOfWarehouseStorage)
	}
}

func (h WarehouseStorageHandler) InsertWarehouseStorage(w http.ResponseWriter, r *http.Request) {
	warehouseStorage := new(models.WarehouseStorage)
	err := json.NewDecoder(r.Body).Decode(&warehouseStorage)
	if err != nil {
		w.Write([]byte("can't decode"))
	}else {
		errUsecase := h.WarehouseStorageUseCase.InsertWarehouseStorage(warehouseStorage)
		if errUsecase != nil {
			fmt.Println(errUsecase)
			w.Write([]byte(fmt.Sprintf("%v",errUsecase)))
		} else {
			byteOfItem, _ := json.Marshal(utils.OtherResponse(http.StatusOK,"Insert Success"))
			w.Header().Set("Content-Type", "application/json")
			w.Write(byteOfItem)
		}
	}
}

func (h WarehouseStorageHandler) UpdateWarehouseStorage(w http.ResponseWriter, r *http.Request) {
	err := h.WarehouseStorageUseCase.UpdateWarehouseStorage(r)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}else {
		byteOfWarehouseStorage, err := json.Marshal(utils.OtherResponse(http.StatusOK, "Update Successfuly"))
		if err != nil {
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(byteOfWarehouseStorage)
	}
}

func (h WarehouseStorageHandler) DeleteWarehouseStorage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := h.WarehouseStorageUseCase.DeleteWarehouseStorage(&id)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}else {
		byteOfWarehouseStorage, err := json.Marshal(utils.OtherResponse(http.StatusOK, "Delete Successfuly"))
		if err != nil {
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(byteOfWarehouseStorage)
	}
}

func WarehouseStorageController(r *mux.Router, model usecases.WarehouseStorageUseCase){
	WarehouseStorageHandler := WarehouseStorageHandler{model}
	sub := r.PathPrefix("").Subrouter()
	sub.Use(middleware.AuthMiddleware)
	sub.HandleFunc("/warehousestorage", WarehouseStorageHandler.ListWarehouseStorage).Methods(http.MethodGet)
	sub.HandleFunc("/warehousestorage", WarehouseStorageHandler.InsertWarehouseStorage).Methods(http.MethodPost)
	sub.HandleFunc("/warehousestorage", WarehouseStorageHandler.UpdateWarehouseStorage).Methods(http.MethodPut)
	sub.HandleFunc("/warehousestorage/{id}", WarehouseStorageHandler.DeleteWarehouseStorage).Methods(http.MethodDelete)
}
