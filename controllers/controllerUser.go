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
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}else {
		errUsecase := h.UserUseCase.InsertUser(user)
		if errUsecase != nil {
			fmt.Println(errUsecase)
			w.Header().Set("Content-Type", "application/json")
			byteOfItem, _ := json.Marshal(utils.OtherResponse(http.StatusNotAcceptable,fmt.Sprintf("%v",errUsecase)))
			w.Write([]byte(byteOfItem))
		} else {
			byteOfItem, _ := json.Marshal(utils.OtherResponse(http.StatusOK,"Insert Success"))
			w.Header().Set("Content-Type", "application/json")
			w.Write(byteOfItem)
		}
	}
}

func (h UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}else {
		errUsecase := h.UserUseCase.LoginUser(user)
		if errUsecase != nil {
			w.Header().Set("Content-Type", "application/json")
			byteOfUser, _ := json.Marshal(utils.OtherResponse(http.StatusNotFound,fmt.Sprintf("%v",errUsecase)))
			w.Write([]byte(byteOfUser))
		} else {
			w.Header().Set("Content-Type", "application/json")
			byteOfUser, _ := json.Marshal(utils.OtherResponse(http.StatusOK,"isi token"))
			w.Write([]byte(byteOfUser))
		}
	}
}

func (h UserHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.UserUseCase.GetUser()
	if err != nil {
		w.Write([]byte(fmt.Sprint("%v",err)))
	}
	byteOfUser, err := json.Marshal(utils.Response(http.StatusOK, "Showing Data successfully", user))
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfUser)
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("can't decode"))
	} else {
		errUsecase := h.UserUseCase.UpdateUser(user)
		if errUsecase != nil {
			fmt.Println(errUsecase)
			w.Write([]byte(fmt.Sprintf("%v", errUsecase)))
		} else {
			byteOfUser, _ := json.Marshal(utils.OtherResponse(http.StatusOK, "Update Success"))
			w.Header().Set("Content-Type", "application/json")
			w.Write(byteOfUser)
		}
	}
}

func UserController(r *mux.Router, model usecases.UserUsecase) {
	UserHandler := UserHandler{model}
	r.HandleFunc("/register", UserHandler.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/login", UserHandler.LoginUser).Methods(http.MethodPost)
	r.HandleFunc("/user", UserHandler.ListUser).Methods(http.MethodGet)
	r.HandleFunc("/user", UserHandler.UpdateUser).Methods(http.MethodPut)
}