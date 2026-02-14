package store

import (
	"errors"
	"sync"
)

type TaskStore struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		data: make(map[string]interface{}),
		mu:   sync.RWMutex{},
	}
}

func (s *TaskStore) AllKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *TaskStore) GetData() map[string]interface{} {
	return s.data
}

func (s *TaskStore) Set(key string, value interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	return nil
}

func (s *TaskStore) Get(key string) (interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, ok := s.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("not found")
}

func (s *TaskStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return nil
}

func (s *TaskStore) Pop() (interface{}, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := s.AllKeys()
	if len(keys) == 0 {
		return nil, false
	}
	key := keys[len(keys)-1]
	if v, ok := s.data[key]; ok {
		delete(s.data, key)
		return v, true
	}
	return nil, false
}
