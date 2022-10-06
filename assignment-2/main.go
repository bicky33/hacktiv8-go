package main

import (
	"assignment-2/config"
	"assignment-2/repository"
	"assignment-2/route"
	"net/http"
)

func main() {
	db := config.GetConnection()
	OrderRepository := repository.NewOrderRepository(db)
	router := route.NewRoutes(OrderRepository)
	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}
	server.ListenAndServe()
}
