package services

import (
	"service/auth/app/repositories"
	"service/auth/app/repositories/dbconnection"
	"service/auth/config"
)

type Services struct {
	UserService  *UserService
	JwtGenerator *JwtGeneratorService
}

func BuildServices(connection *dbconnection.Manager, config config.Config) *Services {
	userRepository := repositories.NewUserRepositrory(connection)
	jwtGenerator := NewJwtGeneratorService(config.Jwt)
	userService := NewUserService(userRepository, jwtGenerator)
	
	return &Services{
		UserService:  userService,
		JwtGenerator: jwtGenerator,
	}
}
