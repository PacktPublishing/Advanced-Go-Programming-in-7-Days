package application

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/4-Database/easy-issues/domain"
)

type IssueService struct {
	IssueRepository  domain.IssueRepository
}

// Returns all the Projects
func (s IssueService) Issues() ([]*domain.Issue, error) {
	return s.IssueRepository.All()
}

// Creates a Project
func (s IssueService) Create(u *domain.Issue) error {
	return s.IssueRepository.Create(u)
}

// Deletes a Project
func (s IssueService) Delete(id int64) error {
	return s.IssueRepository.Delete(id)
}

// Get a Project by id
func (s IssueService) Issue(id int64) (*domain.Issue, error) {
	return s.IssueRepository.GetById(id)
}
