package main

import (
	"fmt"
	"sync"
)

type StoreProducerFunc func() Storer

type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type MemoryStorage struct {
	mu   sync.RWMutex
	data [][]byte // slice of byte slices
}

// func MemoryStoreProducer()

func NewMemoryStore() *MemoryStorage {
	return &MemoryStorage{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStorage) Push(b []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

func (s *MemoryStorage) Fetch(offset int) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset < 0 {
		return nil, fmt.Errorf("offset can not be less than 0")
	}

	if len(s.data) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}
	return s.data[offset], nil
}
