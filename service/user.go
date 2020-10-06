package service

import "github.com/leogsouza/easy-issues/domain"

type service struct {
	repo domain.UserRepository
}

// NewUserService creates an instance of this service
func NewUserService(repo domain.UserRepository) domain.UserService {
	return &service{repo}
}

// Users return all users
func (s *service) Users() ([]*domain.User, error) {
	return s.repo.All()
}

// Create creates an user
func (s *service) Create(u *domain.User) error {
	return s.repo.Create(u)
}

// Delete deletes an user
func (s *service) Delete(id int64) error {
	return s.repo.Delete(id)
}

// User gets an user by id
func (s *service) User(id int64) (*domain.User, error) {
	return s.repo.GetByID(id)
}
