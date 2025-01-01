package jwt

import (
	"service/auth/infrastructure/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockClock struct{}

func (m *mockClock) Now() time.Time {
	time, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	return time
}

func TestGenerate(t *testing.T) {
	// arrange
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjExMzYyMTQzMDUsImlhdCI6MTEzNjIxNDI0NSwiaWQiOiI1In0.BmQ4Nl_JGp8eEIBjtjIAFfwuM4X-qqIaLXAxnudsA6c"
	clock := &mockClock{}
	cfg := config.JwtConfig{
		Secret:     "secret",
		TTLSeconds: 60,
	}

	sut := NewJwtGeneratorService(cfg, clock)
	customClaims := map[string]string{"id": "5"}

	// act
	token, err := sut.Generate(customClaims)

	// assert
	assert.Empty(t, err)
	assert.Equal(t, token, expectedToken)
}
