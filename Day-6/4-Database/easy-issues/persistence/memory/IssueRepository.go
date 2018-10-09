package memory

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/4-Database/easy-issues/domain"
)

const (
	IssuesAllKey = "Issues:all"
	IssueLastId = "Issue:lastId"
)

// IssueRepository concrete implementation of in-memory db
type IssueRepository struct {
	db *cache.Cache
}

func NewIssueRepository() *IssueRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(IssueLastId, int64(0))
	db.SetDefault(IssuesAllKey, []*domain.Issue{})
	return &IssueRepository{
		db: db,
	}
}

func (r IssueRepository)All() ([]*domain.Issue, error)  {
	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		return result.([]*domain.Issue), nil
	} else {
		return nil, errors.New("Empty list")
	}
}

func (r IssueRepository)GetById(id int64) (*domain.Issue, error)  {
	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		items := result.([]*domain.Issue)
		for _, issue := range items {
			if issue.Id == id {
				return issue, nil
			}
		}
		return nil, errors.New("Not Found")
	}
	return nil, errors.New("Not Found")
}

func (r IssueRepository)Create(u *domain.Issue) error  {
	id, _ := r.db.IncrementInt64(IssueLastId, int64(1))
	u.Id = id

	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		result = append(result.([]*domain.Issue), u)
		r.db.Set(IssuesAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r IssueRepository)Delete(id int64) error {
	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		items := result.([]*domain.Issue)
		for i, issue := range items {
			if issue.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(IssuesAllKey, items, cache.NoExpiration)
				return nil
			}
		}
		return errors.New("Not Found")
	}
	return errors.New("Not Found")
}