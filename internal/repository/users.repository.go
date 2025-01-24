package repository

import (
	"anaia-backend/internal/entity"
	"context"
)

const (
	qryInsertUser = `INSERT INTO USERS (first_name, last_name, email, password, role_id) VALUES (?, ?, ?, ?, ?);`

	qryGetUserByEmail = `
		SELECT 
			id_user, 
			first_name, 
			last_name, 
			email, 
			password,
			role_id
		FROM USERS 
		WHERE email = ?;
	`
)

func (r *repo) SaveUser(ctx context.Context, name string, lastname string, email string, password string, roleId int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, name, lastname, email, password, roleId)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}
