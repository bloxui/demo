package store

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Filter string

const (
	FilterAll       Filter = "all"
	FilterActive    Filter = "active"
	FilterCompleted Filter = "completed"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	Priority    Priority
	CreatedAt   time.Time
}

type Stats struct {
	Total     int
	Active    int
	Completed int
}

type Store struct {
	mu     sync.RWMutex
	todos  []*Todo
	nextID uint64
}

func New() *Store {
	s := &Store{}
	s.bootstrap()
	return s
}

func (s *Store) bootstrap() {
	s.todos = []*Todo{
		{
			ID:          "3",
			Title:       "Update documentation",
			Description: "Add new API endpoints to the developer docs",
			Priority:    PriorityLow,
			Completed:   false,
			CreatedAt:   time.Date(2024, time.January, 13, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:          "2",
			Title:       "Review pull requests",
			Description: "Go through the pending PRs and provide feedback",
			Priority:    PriorityMedium,
			Completed:   true,
			CreatedAt:   time.Date(2024, time.January, 14, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:          "1",
			Title:       "Design the new landing page",
			Description: "Create wireframes and mockups for the product landing page",
			Priority:    PriorityHigh,
			Completed:   false,
			CreatedAt:   time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC),
		},
	}
	s.nextID = uint64(len(s.todos) + 1)
	s.sortLocked()
}

func (s *Store) sortLocked() {
	sort.SliceStable(s.todos, func(i, j int) bool {
		return s.todos[i].CreatedAt.After(s.todos[j].CreatedAt)
	})
}

func (s *Store) List(filter Filter) []Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var filtered []Todo
	for _, todo := range s.todos {
		switch filter {
		case FilterActive:
			if todo.Completed {
				continue
			}
		case FilterCompleted:
			if !todo.Completed {
				continue
			}
		}
		filtered = append(filtered, *todo)
	}
	return filtered
}

func (s *Store) Stats() Stats {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stats := Stats{}
	for _, todo := range s.todos {
		stats.Total++
		if todo.Completed {
			stats.Completed++
		} else {
			stats.Active++
		}
	}
	return stats
}

func (s *Store) Add(title, description string, priority Priority) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	trimmedTitle := strings.TrimSpace(title)
	trimmedDescription := strings.TrimSpace(description)

	todo := &Todo{
		ID:          generateID(s.nextID),
		Title:       trimmedTitle,
		Description: trimmedDescription,
		Completed:   false,
		Priority:    priority,
		CreatedAt:   time.Now().UTC(),
	}

	s.nextID++
	s.todos = append([]*Todo{todo}, s.todos...)
	s.sortLocked()

	return *todo
}

func (s *Store) Toggle(id string) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, err := s.findLocked(id)
	if err != nil {
		return Todo{}, err
	}
	todo.Completed = !todo.Completed
	return *todo, nil
}

func (s *Store) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (s *Store) Update(id, title, description string, priority Priority) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, err := s.findLocked(id)
	if err != nil {
		return Todo{}, err
	}

	todo.Title = strings.TrimSpace(title)
	todo.Description = strings.TrimSpace(description)
	todo.Priority = priority
	return *todo, nil
}

func (s *Store) Get(id string) (Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, todo := range s.todos {
		if todo.ID == id {
			return *todo, nil
		}
	}
	return Todo{}, ErrNotFound
}

func (s *Store) findLocked(id string) (*Todo, error) {
	for _, todo := range s.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, ErrNotFound
}

var ErrNotFound = errors.New("todo not found")

func generateID(next uint64) string {
	if next == 0 {
		next = 1
	}
	return fmt.Sprintf("%d", next)
}
