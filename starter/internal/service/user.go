package service

import "github.com/plainkit/starter/internal/domain"

// UserService provides application logic for Users.
type UserService struct {
	Repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService { return &UserService{Repo: repo} }

func (s *UserService) List() ([]domain.User, error) {
	return s.Repo.List()
}

func (s *UserService) Create(name, email string) (domain.User, error) {
	return s.Repo.Create(domain.User{Name: name, Email: email})
}
