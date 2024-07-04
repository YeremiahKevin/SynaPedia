package repository

import (
	// golang package
	"context"
	"fmt"
	"log"
	"time"

	// external package
	"golang.org/x/crypto/bcrypt"
)

// Login is a function to check user credential for login
// it accepts context.Context, string, and string as parameters
// it returns non-empty struct of LoginResponse and nil error when success
// otherwise it returns empty struct of LoginResponse and detailed error
func (repository *Repository) Login(ctx context.Context, username string, password string) (LoginResponse, error) {
	queryResult := db.QueryRowContext(ctx, "SELECT user_id, username, password, full_name FROM \"user\" WHERE username = $1", username)
	var user User
	err := queryResult.Scan(&user.UserID, &user.Username, &user.Password, &user.FullName)
	if err != nil {
		log.Println("Repository Login queryResult.Scan err: ", err)
		return LoginResponse{}, fmt.Errorf("failed to get user data: %s", err.Error())
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

// Register is a function to insert new data to user
// it accepts context.Context and RegisterRequest as parameters
// it returns nil error when success
// otherwise it returns detailed error
func (repository *Repository) Register(ctx context.Context, param RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Repository Register bcrypt.GenerateFromPassword err: ", err)
		return fmt.Errorf("failed to hash password: %s", err.Error())
	}

	_, err = db.ExecContext(ctx, "INSERT INTO \"user\"(username, password, full_name, created_at) VALUES($1, $2, $3, $4)", param.Username, string(hashedPassword), param.FullName, time.Now())
	if err != nil {
		log.Println("Repository Register db.ExecContext err: ", err)
		return fmt.Errorf("failed to insert data: %s", err.Error())
	}

	return nil
}
