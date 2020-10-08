package memory

import (
	"errors"

	"github.com/leogsouza/easy-issues/domain"
	"github.com/patrickmn/go-cache"
)

const (
	ProjectsAllKey = "Projects:all"
	ProjectLastId  = "Project:lastId"
)

type projectRepository struct {
	db *cache.Cache
}

func NewProjectRepository() domain.ProjectRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(ProjectLastId, int64(0))
	db.SetDefault(ProjectsAllKey, []*domain.Project{})
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) GetByID(id int64) (*domain.Project, error) {
	result, ok := r.db.Get(ProjectsAllKey)

	if ok {
		projects := result.([]*domain.Project)
		for _, project := range projects {
			if project.ID == id {
				return project, nil
			}
		}
	}
	return nil, errors.New("Not Found")
}

func (r *projectRepository) All() ([]*domain.Project, error) {
	result, ok := r.db.Get(ProjectsAllKey)

	if ok {
		return result.([]*domain.Project), nil
	}

	return nil, errors.New("Empty list")
}

func (r *projectRepository) Create(p *domain.Project) error {
	id, _ := r.db.IncrementInt64(ProjectLastId, int64(1))
	p.ID = id
	result, ok := r.db.Get(ProjectsAllKey)

	if ok {
		result = append(result.([]*domain.Project), p)
		r.db.Set(ProjectsAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r *projectRepository) Delete(id int64) error {
	result, ok := r.db.Get(ProjectsAllKey)

	if ok {
		projects := result.([]*domain.Project)
		for i, project := range projects {
			if project.ID == id {
				projects = append(projects[:i], projects[i+1:]...)
				r.db.Set(ProjectsAllKey, projects, cache.NoExpiration)
				return nil
			}
		}
	}
	return errors.New("Not Found")
}
