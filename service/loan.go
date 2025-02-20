package loan

import (
	loan "billing-engine/domain"
	"billing-engine/store"
	"errors"
)

type BillingService struct {
	store *store.InMemoryStore
}

func NewBillingService(store *store.InMemoryStore) *BillingService {
	return &BillingService{store: store}
}

func (s *BillingService) CreateLoan(id string, totalWeeks, weeklyPayment int, interest float64) *loan.Loan {
	loan := loan.NewLoan(id, totalWeeks, weeklyPayment, interest)
	s.store.SaveLoan(loan)
	return loan
}

func (s *BillingService) GetLoan(id string) (*loan.Loan, error) {
	loan := s.store.GetLoan(id)
	if loan == nil {
		return nil, errors.New("loan not found")
	}
	return loan, nil
}

func (s *BillingService) GetOutstanding(loanID string) (int, error) {
	loan, err := s.GetLoan(loanID)
	if err != nil {
		return 0, err
	}
	return loan.GetOutstanding(), nil
}

func (s *BillingService) IsDelinquent(loanID string, currentWeek int) (bool, error) {
	loan, err := s.GetLoan(loanID)
	if err != nil {
		return false, err
	}
	return loan.IsDelinquent(currentWeek), nil
}

func (s *BillingService) MakePayment(loanID string, amount int) error {
	loan, err := s.GetLoan(loanID)
	if err != nil {
		return err
	}
	return loan.MakePayment(amount)
}
