package routes

import (
	"net/http"
	"quiz-1/controller"
	"quiz-1/middleware"

	"github.com/gorilla/mux"
)

func HelloRoutes(route *mux.Router) {
	helloWorldControler := middleware.OnlyGet(http.HandlerFunc(controller.HelloWorld))
	route.Handle("/", helloWorldControler)
}
