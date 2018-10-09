package domain

// User represents a user in the System
type User struct {
	Id      int64  `db:"user_id"`
	Name    string `db:"user_name"`
	Surname string `db:"user_surname"`
	Email   string `db:"user_email"`
}

type UserService interface {
	User(id int64) (*User, error)
	Users() ([]*User, error)
	Create(u *User) error
	Delete(id int64) error
}

type UserRepository interface {
	GetById(id int64) (*User, error)
	All() ([]*User, error)
	Create(issue *User) error
	Delete(id int64) error
}
