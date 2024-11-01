package main

import (
	"api/config"
	"api/routes"
	"log"
	"net/http"
)

func main() {
	config.InitDB()
	r := routes.InitRoutes()

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
