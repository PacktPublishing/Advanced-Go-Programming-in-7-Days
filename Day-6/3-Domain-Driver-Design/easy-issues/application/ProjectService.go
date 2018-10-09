package application

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/3-Domain-Driver-Design/easy-issues/domain"
)

type ProjectService struct {
	ProjectRepository  domain.ProjectRepository
}

// Returns all the Projects
func (s *ProjectService) Projects() ([]*domain.Project, error) {
	return s.ProjectRepository.All()
}

// Creates a Project
func (s *ProjectService) Create(u *domain.Project) error {
	return s.ProjectRepository.Create(u)
}

// Deletes a Project
func (s *ProjectService) Delete(id int64) error {
	return s.ProjectRepository.Delete(id)
}

// Get a Project by id
func (s *ProjectService) Project(id int64) (*domain.Project, error) {
	return s.ProjectRepository.GetById(id)
}