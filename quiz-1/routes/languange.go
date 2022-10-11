package routes

import (
	"quiz-1/controller"

	"github.com/gorilla/mux"
)

func LanguageRoutes(router *mux.Router) {
	language := controller.LanguageController{}
	router.HandleFunc("/language", language.Create).Methods("POST")
	router.HandleFunc("/language", language.GetAll).Methods("GET")
	router.HandleFunc("/language/{id}", language.Update).Methods("PATCH")
	router.HandleFunc("/language/{id}", language.Delete).Methods("DELETE")
}
