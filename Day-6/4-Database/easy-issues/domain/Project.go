package domain

// Project represents a top level collection of related issues
type Project struct {
	Id          int64	`db:"project_id"`
	Name        string	`db:"project_name"`
	OwnerId     int64	`db:"project_ownerId"`
	Description string	`db:"project_description"`
}

type ProjectService interface {
	Project(id int64) (*Project, error)
	Projects() ([]*Project, error)
	Create(p *Project) error
	Delete(id int64) error
}

type ProjectRepository interface {
	GetById(id int64) (*Project, error)
	All() ([]*Project, error)
	Create(issue *Project) error
	Delete(id int64) error
}
