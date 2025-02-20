package store

import (
	loan "billing-engine/domain"
	"sync"
)

type InMemoryStore struct {
	loans map[string]*loan.Loan
	mu    sync.Mutex
}

// NewInMemoryStore initializes an in-memory store.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		loans: make(map[string]*loan.Loan),
	}
}

// SaveLoan stores a new loan.
func (s *InMemoryStore) SaveLoan(loan *loan.Loan) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.loans[loan.ID] = loan
}

// GetLoan retrieves a loan by ID.
func (s *InMemoryStore) GetLoan(id string) *loan.Loan {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.loans[id]
}
