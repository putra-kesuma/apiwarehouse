package main

import (
	"apiwarehouse/config"
	"apiwarehouse/middleware"
	"apiwarehouse/routes"
)

func main(){
	//call sql connection db
	db := config.ConnectDb()
	//call function newroute with mux
	route := routes.NewRoute()
	//use middleware
	route.Use(middleware.LoggingMiddleware)
	//call init and send db
	routes.Init(route,db)
	//run server
	routes.RunServer(route)
}
