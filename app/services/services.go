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
	passwordHasher := NewPasswordHasher()
	userService := NewUserService(userRepository, jwtGenerator, passwordHasher)

	return &Services{
		UserService:  userService,
		JwtGenerator: jwtGenerator,
	}
}
