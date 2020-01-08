package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

func (u *UserController) NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/messages", logHandlerCall(u.GetMessages)).Methods("GET")
	router.HandleFunc("/messages", logHandlerCall(u.AddMessage)).Methods("POST", "OPTIONS")

	router.HandleFunc("/categories", logHandlerCall(u.GetCategories)).Methods("GET")

	router.HandleFunc("/forum", logHandlerCall(u.GetCategory)).Methods("GET")
	router.HandleFunc("/forum/subcategories", logHandlerCall(u.GetSubcategories)).Methods("GET")
	router.HandleFunc("/forum/messages", logHandlerCall(u.GetMessages)).Methods("GET")
	router.HandleFunc("/forum/messages", logHandlerCall(u.AddMessage)).Methods("POST", "OPTIONS")

	return router
}

// logHandlerCall logs any handler call
func logHandlerCall(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		log.Printf("Handler function called: %v", name)
		handler(w, r)
	}
}
