package web

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/1-Task-Overview/easy-issues/domain"
)

// Paginator Service for Issue Service
type IssuePaginator struct {
	IssueService domain.IssueService
}

func (p IssuePaginator) Issue(id int64) (*domain.Issue, error)  {
	return p.IssueService.Issue(id)
}

func (p IssuePaginator) Issues(opts *domain.ListOptions) (*domain.ListResponse, error)  {
	return p.IssueService.Issues(opts)
}

// Creates an Issue
func (s IssuePaginator) Create(u *domain.Issue) error {
	return s.IssueService.Create(u)
}

// Deletes an Issue
func (s IssuePaginator) Delete(id int64) error {
	return s.IssueService.Delete(id)
}
