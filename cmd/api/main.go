package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"mustaqel/internal/config"
	"mustaqel/internal/handler"
	"mustaqel/internal/migrations"
	"mustaqel/internal/repository"
	"mustaqel/internal/routes"
	"mustaqel/internal/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run migrations
	if err = migrations.Migrations(db, cfg); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Setup repositories
	userRepo := repository.NewUserRepository(db)

	// Setup services
	userService := service.NewUserService(userRepo)

	// Setup handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup router
	router := routes.Routes(userHandler)

	// Start server
	log.Printf("Server starting on %s", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
