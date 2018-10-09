package domain

// Type Event represents a
type Event struct {
	EventId string
	Type string
}

const CreateUserRegistrationEventType = "CreateUserRegistrationEvent"

// New user Registration event
type CreateUserRegistrationEvent struct {
	Event
	UserEmail string
}

func NewCreateUserRegistrationEvent(id string, email string) CreateUserRegistrationEvent {
	event := new(CreateUserRegistrationEvent)
	event.Type = CreateUserRegistrationEventType
	event.EventId = id
	event.UserEmail = email
	return *event
}
