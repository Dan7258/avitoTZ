package main

import (
	"avito/internal/config"
	"avito/internal/handler"
	"avito/internal/repository"
	"avito/internal/routes"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	config.Init()
	db := &repository.PostgresDB{}
	err := db.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	h := handler.InitHandler(db)
	mux := routes.SetRoutes(h)
	port := ":" + os.Getenv("LOCALHOST_PORT")
	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Starting server on port %s\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
