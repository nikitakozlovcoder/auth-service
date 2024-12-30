package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"service/auth/app/domain/dtos"
	"service/auth/app/repositories/dbconnection"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepositrory struct {
	conn *dbconnection.Manager
}

func NewUserRepositrory(conn *dbconnection.Manager) *UserRepositrory {
	return &UserRepositrory{conn: conn}
}

func (r *UserRepositrory) UserWithHash(ctx context.Context, username string) (dtos.UserWithHash, error) {
	exec := r.conn.Executor(ctx)
	stmt, err := exec.PrepareContext(ctx, "SELECT id, password_hash FROM users WHERE username = $1")
	if err != nil {
		slog.ErrorContext(ctx, err.Error(), slog.Any("error", err))
		return dtos.UserWithHash{}, err
	}

	defer stmt.Close()

	var passwordHash string
	var id int64
	err = stmt.QueryRowContext(ctx, username).Scan(&id, &passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dtos.UserWithHash{}, ErrUserNotFound
		}

		slog.ErrorContext(ctx, err.Error(), slog.Any("error", err))
		return dtos.UserWithHash{}, err

	}

	return dtos.UserWithHash{
		Id:           id,
		PasswordHash: passwordHash,
	}, nil
}
