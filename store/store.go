package store

import (
	"errors"
	"sync"
)

type ReceiptStore struct {
	Receipts map[string]int
	mu       sync.Mutex
}

var Store = ReceiptStore{
	Receipts: make(map[string]int),
}

func (s *ReceiptStore) SaveReceipt(id string, points int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Receipts[id] = points
}

func (s *ReceiptStore) GetPoints(id string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	points, exists := s.Receipts[id]
	if !exists {
		return 0, errors.New("receipt not found")
	}
	return points, nil
}
