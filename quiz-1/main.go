package main

import (
	"net/http"

	"quiz-1/routes"
)

func main() {
	router := routes.MainRoutes()
	server := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}
	server.ListenAndServe()
}
