package domain

// User represents a user in the System
type User struct {
	Id      int64
	Uuid	string
	Name    string
	Surname string
	Email   string
}

type UserService interface {
	Users() ([]*User, error)
	Create(u *User) error
	Update(u *User) error
}

type UserRepository interface {
	All() ([]*User, error)
	Create(u *User) error
	Update(u *User) error
}