package service

import (
	"anaia-backend/internal/models"
	"anaia-backend/internal/repository"
	"context"
)

// Service is the interface implements the business logic of the application.

//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, name string, lastname string, email string, password string, roleId int64) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

// serv is a struct that reference from service to the repository.
type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
