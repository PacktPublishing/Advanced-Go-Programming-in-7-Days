package application

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/domain"
)

type AuthService struct {
	AuthRepository domain.AuthRepository
}

// Returns a User Registration by email
func (s AuthService) GetRegistrationByEmail(email string) (*domain.UserRegistration, error) {
	return s.AuthRepository.GetRegistrationByEmail(email)
}

// Create a new User Registration
func (s AuthService) Create(r *domain.UserRegistration) error {
	return s.AuthRepository.Create(r)
}

