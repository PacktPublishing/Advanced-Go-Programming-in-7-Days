package domain

// Issue represents a Project task
type Issue struct {
	Id          int64    `db:"issue_id"`
	Title       string   `db:"issue_title"`
	Description string   `db:"issue_description"`
	ProjectId   int64    `db:"issue_projectId"`
	OwnerId     int64    `db:"issue_ownerId"`
	Status      Status   `db:"issue_status"`
	Priority    Priority `db:"issue_priority"`
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
