package memory

import (
	"errors"

	"github.com/leogsouza/easy-issues/domain"
	"github.com/patrickmn/go-cache"
)

const (
	IssuesAllKey = "Issues:all"
	IssueLastId  = "Issue:lastId"
)

type issueRepository struct {
	db *cache.Cache
}

func NewIssueRepository() domain.IssueRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(IssueLastId, int64(0))
	db.SetDefault(IssuesAllKey, []*domain.Issue{})

	return &issueRepository{
		db: db,
	}
}

func (r *issueRepository) GetById(id int64) (*domain.Issue, error) {
	result, ok := r.db.Get(IssuesAllKey)

	if ok {
		issues := result.([]*domain.Issue)
		for _, issue := range issues {
			if issue.ID == id {
				return issue, nil
			}
		}
	}
	return nil, errors.New("Not Found")
}

func (r *issueRepository) All() ([]*domain.Issue, error) {
	result, ok := r.db.Get(IssuesAllKey)

	if ok {
		return result.([]*domain.Issue), nil
	}

	return nil, errors.New("Empty list")
}

func (r *issueRepository) Create(issue *domain.Issue) error {
	id, _ := r.db.IncrementInt64(IssueLastId, int64(1))
	issue.ID = id

	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		result = append(result.([]*domain.Issue), issue)
		r.db.Set(IssuesAllKey, result, cache.NoExpiration)
	}
	return nil
}

func (r *issueRepository) Delete(id int64) error {
	result, ok := r.db.Get(IssuesAllKey)
	if ok {
		issues := result.([]*domain.Issue)
		for i, issue := range issues {
			if issue.ID == id {
				issues = append(issues[:i], issues[i+1:]...)
				r.db.Set(IssuesAllKey, issues, cache.NoExpiration)
			}
		}
	}
	return errors.New("Not Found")
}
