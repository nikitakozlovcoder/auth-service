package main

import (
	"log"
	"service/auth/app/repositories/dbconnection"
	"service/auth/app/services"
	"service/auth/config"
	"service/auth/http/handlers"
	"service/auth/http/server"
	"service/auth/logger"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg)
	conn := dbconnection.NewManager(cfg.DbConnectionString)
	defer conn.Close()

	serviceCollection := services.BuildServices(conn, cfg)
	authHandler := handlers.NewAuthHandler(serviceCollection.UserService)
	srv := server.NewServer()
	srv.AddHandler(authHandler)
	err := srv.Run(cfg.Port)

	if err != nil {
		log.Fatal(err)
	}
}
