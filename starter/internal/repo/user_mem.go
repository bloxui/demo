package repo

import "github.com/plainkit/starter/internal/domain"

// InMemoryUserRepo is a simple in-memory repository for demo purposes.
type InMemoryUserRepo struct {
	next  int
	items []domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo { return &InMemoryUserRepo{next: 1} }

func (r *InMemoryUserRepo) List() ([]domain.User, error) {
	// return a copy to avoid external mutation
	out := make([]domain.User, len(r.items))
	copy(out, r.items)
	return out, nil
}

func (r *InMemoryUserRepo) Create(u domain.User) (domain.User, error) {
	u.ID = r.next
	r.next++
	r.items = append(r.items, u)
	return u, nil
}
