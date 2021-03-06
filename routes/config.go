package routes

import (
	"fmt"
	"github.com/gorilla/mux"
//"github.com/gorilla/handlers"
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
	log.Printf("Setting Web Server at host : %v di port : %v", host,port	)
	//c := cors.New(cors.Options{
	//	AllowedOrigins: []string{"http://localhost:8000"},
	//	AllowCredentials: true,
	//})
	//
	//handler := c.Handler(router)
	///
	//headersOk := handlers.AllowedHeaders([]string{"*"})
	//originsOk := handlers.AllowedOrigins([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})


	errListen := http.ListenAndServe(listen, router)
	if errListen != nil {
		log.Fatal(errListen)
	}
}