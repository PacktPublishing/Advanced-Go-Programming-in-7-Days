package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/4-Database/easy-issues/domain"
)

const (
	projectDriver = "sqlite3"
	querySelectAllProjects = "SELECT * FROM projects"
	querySelectProject = "SELECT * FROM projects WHERE project_id=?"
	queryInsertProject = "INSERT INTO projects (project_name, project_ownerId, project_description) VALUES (?, ?, ?)"
	queryDeleteProject = "DELETE FROM projects WHERE project_id=? LIMIT 1"
)

const projectSchema = `CREATE TABLE IF NOT EXISTS projects (
	project_id integer primary key autoincrement,
    project_name text,
	project_ownerId integer,
    project_description text);`

// ProjectRepository concrete implementation of sql db
type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository() *ProjectRepository {
	db, err := sqlx.Open(projectDriver, ":memory:")
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
	r := &ProjectRepository{
		db: db,
	}
	r.Init(projectSchema)
	return r
}

func (r *ProjectRepository)Init(schema string) error  {
	// execute a query on the server
	_, err := r.db.Exec(schema)
	return err
}

func (r *ProjectRepository)All() ([]*domain.Project, error)  {
	rows, err := r.db.Queryx(querySelectAllProjects)
	if err != nil {
		return nil, err
	}
	// iterate over each row
	projects := make([]*domain.Project, 0, 0)
	for rows.Next() {
		var p domain.Project
		err = rows.StructScan(&p)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &p)
	}

	return projects, nil
}

func (r *ProjectRepository)GetById(id int64) (*domain.Project, error)  {
	stmt, err := r.db.Preparex(querySelectProject)
	if err != nil {
		return nil, err
	}

	var p domain.Project
	err = stmt.Get(&p, id)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProjectRepository)Create(p *domain.Project) error {
	stmt, err := r.db.Preparex(queryInsertProject)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(p.Name, p.OwnerId, p.Description)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	p.Id = lastId
	return nil
}

func (r *ProjectRepository)Delete(id int64) error {
	stmt, err := r.db.Preparex(queryDeleteProject)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
