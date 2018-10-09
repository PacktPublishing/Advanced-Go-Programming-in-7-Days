package domain

type RegistrationStatus string

const RegistrationStatusInactive = "RegistrationStatusInactive"
const RegistrationStatusActive = "RegistrationStatusActive"
const RegistrationStatusDeleted = "RegistrationStatusDeleted"

type UserRegistration struct {
	Id           int64              `db:"user_id"`
	Email        string             `db:"user_email"`
	Uuid         string             `db:"user_uuid"`
	Status       RegistrationStatus `db:"user_status"`
	PasswordHash string             `db:"user_password"`
}

type AuthService interface {
	GetRegistrationByEmail(email string) (*UserRegistration, error)
	Create(u *UserRegistration) error
}

type AuthRepository interface {
	GetRegistrationByEmail(email string) (*UserRegistration, error)
	Create(u *UserRegistration) error
}
