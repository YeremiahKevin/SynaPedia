package repository

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type LoginResponse struct {
	IsAllowed bool
}

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	FullName string `db:"full_name"`
}

type RegisterRequest struct {
	Username string
	Password string
	FullName string
}

func (repository *Repository) Login(ctx context.Context, username string, password string) (LoginResponse, error) {
	queryResult := db.QueryRowContext(ctx, "SELECT user_id, username, password, full_name FROM \"user\" WHERE username = $1", username)
	var user User
	err := queryResult.Scan(&user.UserID, &user.Username, &user.Password, &user.FullName)
	if err != nil {
		log.Println("Repository Login queryResult.Scan err: ", err)
		return LoginResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	if user.UserID != 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return LoginResponse{
				IsAllowed: false,
			}, nil
		}

		return LoginResponse{
			IsAllowed: true,
		}, nil
	}

	return LoginResponse{
		IsAllowed: false,
	}, nil
}

func (repository *Repository) Register(ctx context.Context, param RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Repository Register bcrypt.GenerateFromPassword err: ", err)
		return fmt.Errorf("failed to hash password: %w", err)
	}

	_, err = db.ExecContext(ctx, "INSERT INTO \"user\"(username, password, full_name, created_at) VALUES($1, $2, $3, $4)", param.Username, string(hashedPassword), param.FullName, time.Now())
	if err != nil {
		log.Println("Repository Register db.ExecContext err: ", err)
		return fmt.Errorf("failed to insert data: %s", err.Error())
	}

	return nil
}
