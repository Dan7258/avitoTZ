package main

import (
	"avito/internal/config"
	"avito/internal/handler"
	"avito/internal/repository"
	"avito/internal/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	config.Init()
	db := &repository.PostgresDB{}
	err := db.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.InitHandler(db)
	mux := routes.SetRoutes(h)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
