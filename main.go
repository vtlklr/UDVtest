package main

import (
	"UDVtest/config"
	"UDVtest/handlers"
	"UDVtest/repo"
	"UDVtest/service"
)

//docker run --name udv_test -p 5432:5432 -e POSTGRES_USER=udv_test -e POSTGRES_PASSWORD=password -e POSTGRES_DB=udv_test -d postgres:13.3
func main() {

	cfg := config.GetConfig()

	store := repo.New(cfg.Database)

	store.InitDB()
	store.Migrate()

	bookService := service.New(store)
	api := handlers.New(cfg.Web, bookService)
	api.Start()
}
