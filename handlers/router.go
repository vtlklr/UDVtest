package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Api interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	GetBookList(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
	GetHistory(w http.ResponseWriter, r *http.Request)
}

func SetupRoute(api Api) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/book/", api.GetBookList).Methods("GET")
	router.HandleFunc("/api/book/{id}", api.GetBook).Methods("GET")
	router.HandleFunc("/api/book/{id}/items", api.GetItems).Methods("GET")
	router.HandleFunc("/api/book/{id}/history", api.GetHistory).Methods("GET")

	return router
}
