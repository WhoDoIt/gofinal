package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	UserID   int    `json:"user_id"`
	Telegram string `json:"telegram"`
	Email    string `json:"email"`
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Get(ctx context.Context, user_id int) (*User, error) {
	stmt := `SELECT * FROM users WHERE user_id=$1`
	user := &User{}
	err := m.DB.QueryRow(ctx, stmt, user_id).Scan(&user.UserID, &user.Telegram, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
