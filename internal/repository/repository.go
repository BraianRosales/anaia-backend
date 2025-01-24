package repository

import (
	"anaia-backend/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic methods to interact with the database - CRUDS operations.
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, name string, lastname string, email string, password string, roleId int64) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

// repo is a struct that reference to the sqlx.DB to interact with the database. This type is private but can be accessed from the entire repository.
type repo struct {
	db *sqlx.DB
}

// New is a function that initializes the repository.
func New(db *sqlx.DB) Repository {

	return &repo{
		db: db,
	}
}
