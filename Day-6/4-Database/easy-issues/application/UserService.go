package application

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/4-Database/easy-issues/domain"
)

type UserService struct {
	UsersRepository  domain.UserRepository // could abstract it also
}

// All returns all the Users
func (s UserService) Users() ([]*domain.User, error) {
	return s.UsersRepository.All()
}

// All creates a User
func (s UserService) Create(u *domain.User) error {
	return s.UsersRepository.Create(u)
}

// DeleteUser deletes a User
func (s UserService) Delete(id int64) error {
	return s.UsersRepository.Delete(id)
}

// User gets a User by id
func (s UserService) User(id int64) (*domain.User, error) {
	return s.UsersRepository.GetById(id)
}