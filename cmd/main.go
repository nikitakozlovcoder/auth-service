package main

import (
	"log"
	"service/auth/app"
	"service/auth/communications/handlers"
	httpserver "service/auth/communications/http"
	"service/auth/infrastructure/config"
	"service/auth/infrastructure/db/dbconnection"
	"service/auth/infrastructure/logger"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg)
	conn := dbconnection.NewManager(cfg.DbConnectionString)
	defer conn.Close()

	serviceCollection := app.BuildServices(conn, cfg)
	authHandler := handlers.NewAuthHandler(serviceCollection.UserService)
	srv := httpserver.NewServer()
	srv.AddHandler(authHandler)
	err := srv.Run(cfg.Port)

	if err != nil {
		log.Fatal(err)
	}
}
