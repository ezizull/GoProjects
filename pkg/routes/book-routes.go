package route

import (
	"github.com/gorilla/mux"
	"goproject.com/pkg/controller"
)

// RegisterBookRoutes exported
var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
}