package store

import (
	"errors"
	"sync"
)

type AgentStore struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

func NewAgentStore() *AgentStore {
	return &AgentStore{
		data: make(map[string]interface{}),
	}
}

func (s *AgentStore) AllKeys() []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *AgentStore) GetData() map[string]interface{} {
	return s.data
}

func (s *AgentStore) Set(key string, value interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
	return nil
}

func (s *AgentStore) Get(key string) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if v, ok := s.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("not found")
}

func (s *AgentStore) Delete(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
	return nil
}

func (s *AgentStore) Pop() (interface{}, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	keys := s.AllKeys()
	key := keys[len(keys)-1]
	if v, ok := s.data[key]; ok {
		delete(s.data, key)
		return v, true
	}
	return nil, false
}
