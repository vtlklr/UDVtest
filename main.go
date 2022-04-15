package main

import (
	"UDVtest/library"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

/*docker run --name udv_test -p 5432:5432 -e POSTGRES_USER=udv_test -e POSTGRES_PASSWORD=password -e POSTGRES_DB=udv_test -d postgres:13.3*/
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/book/", GetBookList).Methods("GET")
	router.HandleFunc("/api/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/api/book/{id}/items", GetItems).Methods("GET")
	router.HandleFunc("/api/book/{id}/history", GetHistory).Methods("GET")

	//router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
	library.GetDB().DB()
}
