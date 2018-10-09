package domain

// Issue represents a Project task
type Issue struct {
	Id          int64
	Title       string
	Description string
	ProjectId   int64
	OwnerId     int64
	Status      Status
	Priority    Priority
}

type IssueService interface {
	Issue(id int64) (*Issue, error)
	Issues() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}

type IssueRepository interface {
	GetById(id int64) (*Issue, error)
	All() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}
