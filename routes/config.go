package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func NewRoute() *mux.Router {
	route := mux.NewRouter()
	return route
}

func RunServer(router *mux.Router) {
	// load .env file
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal(errEnv)
	}

	fmt.Println("Your Server Is Running! Awesome Brayy")
	host := os.Getenv("SERVER_IP_HOST")
	port := os.Getenv("SERVER_PORT")
	listen := fmt.Sprintf("%v:%v",host,port)
	fmt.Println("Setting Web Server at host : " + host)
	fmt.Println("Setting Web Server at port : " + port)
	errListen := http.ListenAndServe(listen, router)
	if errListen != nil {
		log.Fatal(errListen)
	}
}