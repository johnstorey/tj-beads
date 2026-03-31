package db

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Username string
	Password string
}

func (d *DB) CreateUserTable(ctx context.Context) error {
	_, err := d.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		)
	`)
	return err
}

func (d *DB) CreateUser(ctx context.Context, username, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result, err := d.ExecContext(ctx,
		"INSERT INTO users (username, password) VALUES (?, ?)",
		username, string(hashedPassword),
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Username: username,
		Password: string(hashedPassword),
	}, nil
}

func (d *DB) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := d.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *DB) CheckPassword(username, password string) bool {
	user, err := d.GetUserByUsername(context.Background(), username)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}