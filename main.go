package main

import (
	"log"
	"mustaqel/config"
	routes "mustaqel/routes"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func getPort() string {
	if value := os.Getenv("APP_PORT"); value != "" {
		return value
	}
	return "8080"
}

func main() {
	config.ConnectDB()

	port := getPort()

	router := chi.NewRouter()
	web := routes.SetupWebRoutes()
	api := routes.SetupApiRoutes()
	router.Mount("/", web)
	router.Mount("/api", api)
	log.Fatal(http.ListenAndServe(":"+port, router))
	log.Printf("Server Started :%s", port)
}
