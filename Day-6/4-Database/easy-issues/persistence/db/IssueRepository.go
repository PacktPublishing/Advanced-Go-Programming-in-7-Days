package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/4-Database/easy-issues/domain"
)

const (
	issuesDriver = "sqlite3"
	querySelectAllIssues = "SELECT * FROM issues"
	querySelectIssue = "SELECT * FROM issues WHERE issue_id=?"
	queryInsertIssue = "INSERT INTO issues (issue_title, issue_description, issue_projectId, issue_ownerId, issue_status, issue_priority) VALUES (?, ?, ?, ?, ?, ?)"
	queryDeleteIssue = "DELETE FROM issues WHERE issue_id=? LIMIT 1"
)

const issueSchema = `CREATE TABLE IF NOT EXISTS issues (
	issue_id integer primary key autoincrement,
    issue_title text,
	issue_description text,
	issue_projectId integer,
	issue_ownerId integer,
	issue_status text,
    issue_priority text);`


// ProjectRepository concrete implementation of sql db
type IssueRepository struct {
	db *sqlx.DB
}

func NewIssueRepository() *IssueRepository {
	db, err := sqlx.Open(issuesDriver, ":memory:")
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	r := &IssueRepository{
		db: db,
	}
	r.Init(issueSchema)
	return r
}

func (r *IssueRepository)Init(schema string) error  {
	// execute a query on the server
	_, err := r.db.Exec(schema)
	return err
}

func (r *IssueRepository)All()([]*domain.Issue, error)  {
	rows, err := r.db.Queryx(querySelectAllIssues)
	if err != nil {
		return nil, err
	}

	issues := make([]*domain.Issue, 0, 0)
	for rows.Next() {
		var u domain.Issue
		err = rows.StructScan(&u)
		if err != nil {
			return nil, err
		}
		issues = append(issues, &u)
	}

	return issues, nil
}

func (r *IssueRepository)GetById(id int64)(*domain.Issue, error)  {
	stmt, err := r.db.Preparex(querySelectIssue)
	if err != nil {
		return nil, err
	}

	var u domain.Issue
	err = stmt.Get(&u, id)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *IssueRepository)Create(u *domain.Issue) error  {
	stmt, err := r.db.Preparex(queryInsertIssue)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Title, u.Description, u.OwnerId, u.ProjectId, u.Status, u.Priority)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = lastId
	return nil
}

func (r *IssueRepository)Delete(id int64) error  {
	stmt, err := r.db.Preparex(queryDeleteIssue)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
