package application

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/domain"
	"github.com/Shopify/sarama"
	"fmt"
	"reflect"
	"encoding/json"
)

type UserService struct {
	UsersRepository  domain.UserRepository
	EventsConsumer sarama.Consumer
}

// Returns all the Users
func (s UserService) Users() ([]*domain.User, error) {
	return s.UsersRepository.All()
}

// Creates a User
func (s UserService) Create(u *domain.User) error {
	return s.UsersRepository.Create(u)
}

// Updates a User
func (s UserService) Update(u *domain.User) error {
	return s.UsersRepository.Update(u)
}

func (s UserService) OnEventTypeHandle(msgVal []byte, eventType interface{})  {
	var err error
	fmt.Println(reflect.ValueOf(eventType))
	switch eventType {
	case domain.CreateUserRegistrationEventType:
		event := new(domain.CreateUserRegistrationEvent)
		err = json.Unmarshal(msgVal, &event)
		if err != nil {
			fmt.Printf("Error parsing event %v", msgVal)
		}
		s.UsersRepository.Create(
			&domain.User{
				Uuid: event.EventId,
				Email: event.UserEmail,
			},
		)
	default:
		fmt.Println("Unknown Event type: ", eventType)
	}
}