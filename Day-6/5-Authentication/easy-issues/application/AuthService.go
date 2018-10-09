package application

import "github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues/domain"

type AuthService struct {
	AuthRepository domain.AuthRepository
}

// Returns all the Projects
func (s AuthService) GetRegistrationByEmail(email string) (*domain.UserRegistration, error) {
	return s.AuthRepository.GetRegistrationByEmail(email)
}
