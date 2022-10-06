package route

import (
	"assignment-2/controller"
	"assignment-2/repository"

	"github.com/gorilla/mux"
)

func NewRoutes(OrderRepository repository.OrderRepository) *mux.Router {
	orderController := controller.OrderController{OrderRepository: OrderRepository}
	router := mux.NewRouter()
	router.HandleFunc("/orders", orderController.Create).Methods("POST")
	router.HandleFunc("/orders", orderController.GetAll).Methods("GET")
	router.HandleFunc("/orders/{id}", orderController.Delete).Methods("DELETE")
	router.HandleFunc("/orders/{id}", orderController.Update).Methods("PUT")
	return router
}
