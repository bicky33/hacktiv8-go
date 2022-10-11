package routes

import "github.com/gorilla/mux"

func MainRoutes() *mux.Router {
	router := mux.NewRouter()
	HelloRoutes(router)
	LanguageRoutes(router)
	PalindormRoutes(router)
	return router
}
