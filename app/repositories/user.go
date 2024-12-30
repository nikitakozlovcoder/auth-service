package repositories

import (
	"context"
	"log/slog"
	"service/auth/app/repositories/dbconnection"
)

type UserRepositrory struct {
	conn *dbconnection.Manager
}

func NewUserRepositrory(conn *dbconnection.Manager) *UserRepositrory {
	return &UserRepositrory{conn: conn}
}

func (r *UserRepositrory) UserExists(ctx context.Context, username, passwordHash string) (bool, error) {
	exec := r.conn.Executor(ctx)
	var exists bool
	stmt, err := exec.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 AND password_hash = $2)")
	if err != nil {
		slog.ErrorContext(ctx, "Error preparing statement", slog.Any("error", err))
		return false, err
	}

	stmt.QueryRowContext(ctx, username, passwordHash).Scan(&exists)
	return exists, nil
}
