package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/leogsouza/easy-issues/domain"
)

const (
	issuesDriver         = "sqlite3"
	querySelectAllIssues = "SELECT * FROM issues"
	querySelectIssue     = "SELECT * FROM issues WHERE id = ?"
	queryInsertIssue     = "INSERT INTO issues (title, description, project_id, owner_id, status, priority) VALUES (?, ?, ?, ?, ?, ?)"
	queryDeleteIssue     = "DELETE FROM issues WHERE id = ? LIMIT 1"
)

const issueSchema = `CREATE TABLE IF NOT EXISTS issues (
		id integer primary key autoincrement,
		title text,
		description text,
		project_id integer,
		owner_id integer,
		status text,
		priority text);`

type issueRepository struct {
	db *sqlx.DB
}

func NewIssueRepository() domain.IssueRepository {
	db, err := sqlx.Open(issuesDriver, ":memory:")
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	r := &issueRepository{
		db: db,
	}

	r.Init(issueSchema)

	return r
}

func (r *issueRepository) Init(schema string) error {
	_, err := r.db.Exec(schema)
	return err
}

func (r *issueRepository) GetById(id int64) (*domain.Issue, error) {
	stmt, err := r.db.Preparex(querySelectIssue)

	if err != nil {
		return nil, err
	}

	var issue domain.Issue
	err = stmt.Get(&issue, id)

	if err != nil {
		return nil, err
	}

	return &issue, nil
}

func (r *issueRepository) All() ([]*domain.Issue, error) {
	rows, err := r.db.Queryx(querySelectAllIssues)

	if err != nil {
		return nil, err
	}

	issues := make([]*domain.Issue, 0, 0)

	for rows.Next() {
		var issue domain.Issue
		err := rows.StructScan(&issue)
		if err != nil {
			return nil, err
		}

		issues = append(issues, &issue)
	}

	return issues, nil
}

func (r *issueRepository) Create(issue *domain.Issue) error {
	stmt, err := r.db.Preparex(queryInsertIssue)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(issue.Title, issue.Description, issue.OwnerID, issue.ProjectID, issue.Status, issue.Priority)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	issue.ID = lastId

	return nil
}

func (r *issueRepository) Delete(id int64) error {
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
