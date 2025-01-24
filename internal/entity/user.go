package entity

type User struct {
	ID       int64  `db:"id_user"`
	Name     string `db:"first_name"`
	LastName string `db:"last_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	RoleId   int64  `db:"role_id"`
}
