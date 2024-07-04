package usecase

import (
	// golang package
	"context"
	"log"

	// internal package
	"SynaPedia/repository"
)

// Login is a function to check user credential for login
// it accepts context.Context, string, and string as parameters
// it returns non-empty struct of LoginResponse and nil error when success
// otherwise it returns empty struct of LoginResponse and detailed error
func (usecase *Usecase) Login(ctx context.Context, username string, password string) (LoginResponse, error) {
	result, err := usecase.Repository.Login(ctx, username, password)
	if err != nil {
		log.Println("Usecase Login err: ", err)
		return LoginResponse{}, err
	}

	return LoginResponse(result), nil
}

// Register is a function to register new user
// it accepts context.Context and RegisterRequest as parameters
// it returns nil error when success
// otherwise it returns detailed error
func (usecase *Usecase) Register(ctx context.Context, param RegisterRequest) error {
	err := usecase.Repository.Register(ctx, repository.RegisterRequest(param))
	if err != nil {
		log.Println("Usecase Register err: ", err)
		return err
	}

	return nil
}
