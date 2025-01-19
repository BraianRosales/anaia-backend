package models

type User struct {
	ID       int64  `json:"id_user"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}
