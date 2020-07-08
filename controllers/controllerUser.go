package controllers

import (
	"apiwarehouse/models"
	"apiwarehouse/usecases"
	"apiwarehouse/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	UserUseCase usecases.UserUsecase
}

func (h UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	//err := json.NewDecoder(r.Body).Decode(&user)
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	//if err != nil {
	//	w.Write([]byte("can't decode"))
	//}else {
		errUsecase := h.UserUseCase.InsertUser(user)
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
	//}
}

func UserController(r *mux.Router, model usecases.UserUsecase) {
	UserHandler := UserHandler{model}
	r.HandleFunc("/register", UserHandler.RegisterUser).Methods(http.MethodPost)
}