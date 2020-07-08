package controllers

import (
	"apiwarehouse/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)


func AuthController(r *mux.Router){

	r.HandleFunc("/auth", Login).Methods(http.MethodPost)
	r.HandleFunc("/token", ValidasiToken).Methods(http.MethodPost)

}

func ValidasiToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if len(token) == 0 {
		token = ""
	} else {
		tokenVal := strings.Split(token, "Bearer ")
		token = strings.Trim(tokenVal[1], "")
		fmt.Println(token)
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("can't decode"))
	}else {
		if user.Username == "putra" && user.Password == "inipass"{
			sign:= jwt.New(jwt.GetSigningMethod("HS256"))
			token, _ := sign.SignedString([]byte("secret"))
			w.Write([]byte(token))
		}
	}
}
