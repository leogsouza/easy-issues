package domain

// Project represents a top level collection of related issues
type Project struct {
	ID          int64
	Name        string
	OwnerID     int64
	Description string
}

// ProjectService is an interface to handle projects
type ProjectService interface {
	Project(id int64) (*Project, error)
	Projects() ([]*Project, error)
	Create(p *Project) error
	Delete(id int64) error
}

// ProjectRepository is a repository which interact with the database layer
type ProjectRepository interface {
	GetByID(id int64) (*Project, error)
	All() ([]*Project, error)
	Create(p *Project) error
	Delete(id int64) error
}
