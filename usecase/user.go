package usecase

import (
	"SynaPedia/repository"
	"context"
	"log"
)

func (usecase *Usecase) Login(ctx context.Context, username string, password string) (LoginResponse, error) {
	result, err := usecase.Repository.Login(ctx, username, password)
	if err != nil {
		log.Println("Usecase Login err: ", err)
		return LoginResponse{}, err
	}

	return LoginResponse(result), nil
}

func (usecase *Usecase) Register(ctx context.Context, param RegisterRequest) error {
	err := usecase.Repository.Register(ctx, repository.RegisterRequest(param))
	if err != nil {
		log.Println("Usecase Register err: ", err)
		return err
	}

	return nil
}
