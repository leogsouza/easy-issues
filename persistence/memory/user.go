package memory

import (
	"errors"

	"github.com/leogsouza/easy-issues/domain"
	"github.com/patrickmn/go-cache"
)

const (
	UsersAllKey = "Users:all"
	UserLastId  = "User:lastId"
)

// userRepository concrete implementation of in-memory db
type userRepository struct {
	db *cache.Cache
}

func NewUserRepository() domain.UserRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(UserLastId, int64(0))
	db.SetDefault(UsersAllKey, []*domain.User{})

	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByID(id int64) (*domain.User, error) {

	result, ok := r.db.Get(UserLastId)

	if ok {
		items := result.([]*domain.User)
		for _, user := range items {
			if user.ID == id {
				return user, nil
			}
		}

	}
	return nil, errors.New("Empty")
}

func (r *userRepository) All() ([]*domain.User, error) {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		return result.([]*domain.User), nil
	}

	return nil, errors.New("Empty")
}

func (r *userRepository) Create(u *domain.User) error {
	id, _ := r.db.IncrementInt64(UserLastId, int64(1))
	u.ID = id

	result, ok := r.db.Get(UsersAllKey)
	if ok {
		result = append(result.([]*domain.User), u)
		r.db.Set(UsersAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r *userRepository) Delete(id int64) error {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		items := result.([]*domain.User)
		for i, user := range items {
			if user.ID == id {
				items = append(items[:i], items[i+1:]...)
				return nil
			}
		}
		return errors.New("Not found")
	}
	return errors.New("Not found")
}
