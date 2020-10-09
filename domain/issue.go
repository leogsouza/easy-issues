package domain

// Issue represents a Project task
type Issue struct {
	ID          int64    `db:"id"`
	Title       string   `db:"title"`
	Description string   `db:"description"`
	ProjectID   int64    `db:"project_id"`
	OwnerID     int64    `db:"owner_id"`
	Status      Status   `db:"status"`
	Priority    Priority `db:"priority"`
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
