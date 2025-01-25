package service

import (
	"anaia-backend/encryption"
	"anaia-backend/internal/models"
	"context"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidPassword   = errors.New("invalid password")
)

func (s *serv) RegisterUser(ctx context.Context, name string, lastname string, email string, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	//pp = array of bytes | pp is the references of the arrays in go.
	pp, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(pp)

	return s.repo.SaveUser(ctx, name, lastname, email, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	aob, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	descryptedPassword, err := encryption.Decrypt(aob)
	if err != nil {
		return nil, err
	}

	if string(descryptedPassword) != password {
		return nil, ErrInvalidPassword
	}

	return &models.User{
		ID:       u.ID,
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
	}, nil
}
