package application

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/1-Task-Overview/easy-issues/domain"
)

type IssueService struct {
	IssueRepository  domain.IssueRepository
}

// Returns all the Issues
func (s IssueService) Issues(opts *domain.ListOptions) (*domain.ListResponse, error) {
	return s.IssueRepository.All(opts)
}

// Get an Issue by id
func (s IssueService) Issue(id int64) (*domain.Issue, error) {
	return s.IssueRepository.GetById(id)
}

// Creates an Issue
func (s IssueService) Create(u *domain.Issue) error {
	return s.IssueRepository.Create(u)
}

// Deletes an Issue
func (s IssueService) Delete(id int64) error {
	return s.IssueRepository.Delete(id)
}

