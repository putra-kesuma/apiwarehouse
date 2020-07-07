package controllers

import (
	"apiwarehouse/usecases"
	"apiwarehouse/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type WarehouseHandler struct {
	WarehouseUseCase usecases.WarehouseUseCase
}

func (h WarehouseHandler) ListWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouse, err := h.WarehouseUseCase.GetWarehouse()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfWarehouse, err := json.Marshal(
		utils.Response(http.StatusOK,"Showing Data successfully", warehouse))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfWarehouse)
}

func (h WarehouseHandler) InsertWarehouse(w http.ResponseWriter, r *http.Request) {
	err := h.WarehouseUseCase.InsertWarehouse(r)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfWarehouse,err := json.Marshal(utils.OtherResponse(http.StatusOK,"Insert Successfuly"))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfWarehouse)
}

func (h WarehouseHandler) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	err := h.WarehouseUseCase.UpdateWarehouse(r)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfWarehouse,err := json.Marshal(utils.OtherResponse(http.StatusOK,"Update Successfuly"))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfWarehouse)
}

func (h WarehouseHandler) DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := h.WarehouseUseCase.DeleteWarehouse(&id)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfWarehouse,err := json.Marshal(utils.OtherResponse(http.StatusOK,"Delete Successfuly"))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfWarehouse)
}

func WarehouseController(r *mux.Router, model usecases.WarehouseUseCase){
	WarehouseHandler := WarehouseHandler{model}
	r.HandleFunc("/warehouse", WarehouseHandler.ListWarehouse).Methods(http.MethodGet)
	r.HandleFunc("/warehouse", WarehouseHandler.InsertWarehouse).Methods(http.MethodPost)
	r.HandleFunc("/warehouse", WarehouseHandler.UpdateWarehouse).Methods(http.MethodPut)
	r.HandleFunc("/warehouse/{id}", WarehouseHandler.DeleteWarehouse).Methods(http.MethodDelete)
}