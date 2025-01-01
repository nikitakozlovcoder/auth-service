package app

import (
	"service/auth/app/jwt"
	"service/auth/app/password"
	"service/auth/app/users"
	"service/auth/infrastructure/clock"
	"service/auth/infrastructure/config"
	"service/auth/infrastructure/db/dbconnection"
)

type Services struct {
	UserService  *users.UserService
	JwtGenerator *jwt.JwtGeneratorService
}

func BuildServices(connection *dbconnection.Manager, config config.Config) *Services {
	timeClock := clock.NewTimeClock()
	userRepository := users.NewUserRepositrory(connection)
	jwtGenerator := jwt.NewJwtGeneratorService(config.Jwt, timeClock)
	passwordHasher := password.NewPasswordComparer()
	userService := users.NewUserService(userRepository, jwtGenerator, passwordHasher)

	return &Services{
		UserService:  userService,
		JwtGenerator: jwtGenerator,
	}
}
