package recipes

import (
	"errors"
	"fmt"
)

type MemStore struct {
	storage map[string]Recipe
}

func NewMemStore() *MemStore {
	return &MemStore{storage: map[string]Recipe{}}
}

func (s *MemStore) Add(name string, recipe Recipe) error {
	fmt.Println("Add")
	_, exists := s.storage[name]

	if exists {
		return errors.New("Recipe already exists")
	}

	s.storage[name] = recipe
	return nil
}

func (s *MemStore) Get(name string) (Recipe, error) {
	fmt.Println("Get")

	val, exists := s.storage[name]

	if !exists {
		return val, errors.New("Recipe doesn't exists")
	}

	return val, nil
}

func (s *MemStore) Update(name string, recipe Recipe) error {
	_, exists := s.storage[name]

	if !exists {
		return errors.New("Recipe doesn't exists")
	}

	s.storage[name] = recipe

	return nil
}

func (s *MemStore) List() (map[string]Recipe, error) {
	fmt.Println("List")
	return s.storage, nil
}

func (s *MemStore) Remove(name string) error {
	_, exists := s.storage[name]

	if !exists {
		return errors.New("Recipe doesn't exists")
	}

	delete(s.storage, name)

	return nil
}
