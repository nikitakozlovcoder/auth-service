package services

import (
	"service/auth/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtGeneratorService struct {
	config config.JwtConfig
}

func NewJwtGeneratorService(config config.JwtConfig) *JwtGeneratorService {
	return &JwtGeneratorService{config: config}
}

func (s *JwtGeneratorService) Generate(claims map[string]string) (string, error) {
	jwtClaims := jwt.MapClaims{}
	for k, v := range claims {
		jwtClaims[k] = v
	}

	jwtClaims["iat"] = time.Now().Unix()
	jwtClaims["exp"] = time.Now().Add(time.Duration(s.config.TTLSeconds) * time.Second).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(s.config.Secret))
}
