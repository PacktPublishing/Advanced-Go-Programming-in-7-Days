package memory

import (
	"github.com/patrickmn/go-cache"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/domain"
	"errors"
)

const (
	UsersAllKey = "Users:all"
)

// UserRepository concrete implementation of in-memory db
type UserRepository struct {
	// Need locks here
	db *cache.Cache
}

func NewUserRepository() *UserRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(UsersAllKey, []*domain.User{})
	return &UserRepository{
		db: db,
	}
}

func (r * UserRepository)All() ([]*domain.User, error)  {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		return result.([]*domain.User), nil
	} else {
		return nil, errors.New("Empty")
	}
}

func (r * UserRepository)GetById(id string) (*domain.User, error)  {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		items := result.([]*domain.User)
		for _, user := range items {
			if user.Uuid == id {
				return user, nil
			}
		}
		return nil, errors.New("Empty")
	}
	return nil, errors.New("Empty")
}

func (r * UserRepository)Create(u *domain.User) error  {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		result = append(result.([]*domain.User), u)
		r.db.Set(UsersAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r * UserRepository)Update(u *domain.User) error  {
	result, ok := r.db.Get(UsersAllKey)
	items := result.([]*domain.User)
	var idx int
	if ok {
		for i, user := range items {
			if user.Uuid == u.Uuid {
				idx = i
				break
			}
		}
		return errors.New("Not Found")
	}

	if idx >= 0 {
		items = append(items[:idx], items[idx+1:]...)
		items = append(items, u)
		r.db.Set(UsersAllKey, result, cache.NoExpiration)
	}
	return nil
}
