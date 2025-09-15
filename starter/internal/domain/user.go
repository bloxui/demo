package domain

// User is a minimal example domain entity.
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository abstracts persistence for Users.
type UserRepository interface {
	List() ([]User, error)
	Create(User) (User, error)
}
