package password

import (
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type PasswordComparerService struct{}

func NewPasswordComparer() *PasswordComparerService {
	return &PasswordComparerService{}
}

func (*PasswordComparerService) Compare(passwordHash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		if !errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			slog.Error(err.Error(), slog.Any("error", err))
			return false, err
		}

		return false, nil
	}

	return true, nil
}
