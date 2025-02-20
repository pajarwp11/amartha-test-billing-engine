package main

import "errors"

func (s *LoanStore) SaveLoan(loan *Loan) {
	s.Loans[loan.ID] = loan
}

func (s *LoanStore) GetLoan(id string) (*Loan, error) {
	loan, exists := s.Loans[id]
	if !exists {
		return nil, errors.New("loan not found")
	}
	return loan, nil
}
