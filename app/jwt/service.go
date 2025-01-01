package jwt

import (
	"service/auth/infrastructure/clock"
	"service/auth/infrastructure/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtGeneratorService struct {
	config config.JwtConfig
	clock  clock.Clock
}

func NewJwtGeneratorService(config config.JwtConfig, clock clock.Clock) *JwtGeneratorService {
	return &JwtGeneratorService{config: config, clock: clock}
}

func (s *JwtGeneratorService) Generate(claims map[string]string) (string, error) {
	jwtClaims := jwt.MapClaims{}
	for k, v := range claims {
		jwtClaims[k] = v
	}

	jwtClaims["iat"] = s.clock.Now().Unix()
	jwtClaims["exp"] = s.clock.Now().Add(time.Duration(s.config.TTLSeconds) * time.Second).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(s.config.Secret))
}
