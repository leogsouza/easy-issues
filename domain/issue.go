package domain

// Issue represents a Project task
type Issue struct {
	ID        int64
	Title     string
	ProjectID int64
	OwnerID   int64
	Status    Status
	Priority  Priority
}

// IssueService contains methods do Handle issue model
type IssueService interface {
	Issue(id int64) (*Issue, error)
	Issues() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}

// IssueRepository contains methods to handles issue database operations
type IssueRepository interface {
	GetById(id int64) (*Issue, error)
	All() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}
