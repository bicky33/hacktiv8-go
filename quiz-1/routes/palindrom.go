package routes

import (
	"quiz-1/controller"

	"github.com/gorilla/mux"
)

func PalindormRoutes(router *mux.Router) {
	router.HandleFunc("/palindrom", controller.PalindromController)
}
