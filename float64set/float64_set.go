// Code generated by internal/gen.tmpl, DO NOT EDIT.

package float64set

import (
	"sync"
)

type set struct {
	mu   sync.RWMutex
	data map[float64]struct{}
}

// New create and return instance of float64 set
func New(float64s ...float64) *set {
	s := NewWithSize(len(float64s))
	s.Adds(float64s...)
	return s
}

// New create and return instance of float64 set
func NewWithSize(size int) *set {
	return &set{
		data: make(map[float64]struct{}, size),
	}
}

// Add item to the set.
func (s *set) Add(item float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[item] = struct{}{}
}

// Add items to the set.
func (s *set) Adds(items ...float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range items {
		s.data[item] = struct{}{}
	}
}

// Remove item from the set.
func (s *set) Remove(item float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, item)
}

// Remove items from the set.
func (s *set) Removes(items ...float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range items {
		delete(s.data, item)
	}
}

// Exist determines whether the item is exists in the set.
func (s *set) Exist(item float64) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[item]
	return ok
}

// List convert set to an array.
func (s *set) List(item float64) []float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	i := 0
	l := make([]float64, len(s.data))
	for val := range s.data {
		l[i] = val
		i++
	}

	return l
}
