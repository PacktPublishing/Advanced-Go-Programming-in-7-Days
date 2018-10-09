package memory

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/3-Domain-Driver-Design/easy-issues/domain"
)

const (
	ProjectsAllKey = "Projects:all"
	ProjectLastId = "Project:lastId"
)


// ProjectRepository concrete implementation of in-memory db
type ProjectRepository struct {
	db *cache.Cache
}

func NewProjectRepository() *ProjectRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(ProjectLastId, int64(0))
	db.SetDefault(ProjectsAllKey, []*domain.Project{})
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository)All() ([]*domain.Project, error)  {
	result, ok := r.db.Get(ProjectsAllKey)
	if ok {
		return result.([]*domain.Project), nil
	} else {
		return nil, errors.New("Empty list")
	}
}

func (r *ProjectRepository)GetById(id int64) (*domain.Project, error)  {
	result, ok := r.db.Get(ProjectsAllKey)
	if ok {
		items := result.([]*domain.Project)
		for _, project := range items {
			if project.Id == id {
				return project, nil
			}
		}
		return nil, errors.New("Not Found")
	}
	return nil, errors.New("Not Found")
}

func (r *ProjectRepository)Create(p *domain.Project) error  {
	id, _ := r.db.IncrementInt64(ProjectLastId, int64(1))
	p.Id = id

	result, ok := r.db.Get(ProjectsAllKey)
	if ok {
		result = append(result.([]*domain.Project), p)
		r.db.Set(ProjectsAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r *ProjectRepository)Delete(id int64) error  {
	result, ok := r.db.Get(ProjectsAllKey)
	if ok {
		items := result.([]*domain.Project)
		for i, project := range items {
			if project.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(ProjectsAllKey, items, cache.NoExpiration)
				return nil
			}
		}
		return errors.New("Not Found")
	}
	return errors.New("Not Found")
}