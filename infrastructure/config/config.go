package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	Env                string
	DbConnectionString string
	Jwt                JwtConfig
}

type JwtConfig struct {
	Secret     string
	TTLSeconds int64
}

const (
	EnvLocal       = "local"
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func Load() Config {
	env := os.Getenv("ENVIRONMENT")

	if env == "" {
		env = EnvLocal
		godotenv.Load(".env.local")
	} else {
		err := godotenv.Load(".env." + env)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	err := godotenv.Load()
	if err != nil && env == EnvLocal {
		log.Fatal("Error loading .env file")
	}

	jwtTTL, err := strconv.ParseInt(os.Getenv("JWT_TTL_SECONDS"), 10, 64)
	if err != nil {
		log.Fatal("Cant parse JWT_TTL_SECONDS ", err)
	}

	return Config{
		Port: os.Getenv("PORT"),
		Env:  env,
		DbConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		Jwt:  JwtConfig{Secret: os.Getenv("JWT_SECRET"), TTLSeconds: jwtTTL},
	}
}
