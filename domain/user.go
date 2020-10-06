package domain

// User represents a user in the System
type User struct {
	ID      int64
	Name    string
	Surname string
	Email   string
}

// UserService is an inteface which interacts with the repository
type UserService interface {
	User(id int64) (*User, error)
	Users() ([]*User, error)
	Create(u *User) error
	Delete(id int64) error
}

// UserRepository is an interface which interacts with the database layer
type UserRepository interface {
	GetByID(id int64) (*User, error)
	All() ([]*User, error)
	Create(u *User) error
	Delete(id int64) error
}
