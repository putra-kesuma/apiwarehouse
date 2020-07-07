package controllers

import (
	"apiwarehouse/usecases"
	"apiwarehouse/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type ReportWarehouseHandler struct {
	ReportWarehouseUseCase usecases.ReportWarehouseUseCase
}

func (h ReportWarehouseHandler) ReportWarehouse(w http.ResponseWriter, r *http.Request) {
	reportWarehouse, err := h.ReportWarehouseUseCase.GetReportWarehouse()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfReportWarehouse, err := json.Marshal(
		utils.Response(http.StatusOK,"Showing Data successfully", reportWarehouse))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfReportWarehouse)
}

func ReportWarehouseController(r *mux.Router, model usecases.ReportWarehouseUseCase){
	ReportWarehouseHandler := ReportWarehouseHandler{model}
	r.HandleFunc("/report", ReportWarehouseHandler.ReportWarehouse).Methods(http.MethodGet)
}
