package loan

import (
	"billing-engine/store"
	"testing"
)

func TestBillingService(t *testing.T) {
	store := store.NewInMemoryStore()
	billingService := NewBillingService(store)

	loan := billingService.CreateLoan("1", 50, 110000, 10.0)
	if loan.ID != "1" {
		t.Errorf("Expected loan ID to be '1', got %s", loan.ID)
	}
}

func TestGetOutstanding(t *testing.T) {
	store := store.NewInMemoryStore()
	billingService := NewBillingService(store)

	billingService.CreateLoan("1", 50, 110000, 10.0)
	outstanding, err := billingService.GetOutstanding("1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if outstanding != 5500000 {
		t.Errorf("Expected outstanding amount to be 5500000, got %d", outstanding)
	}
}

func TestMakePayment(t *testing.T) {
	store := store.NewInMemoryStore()
	billingService := NewBillingService(store)

	billingService.CreateLoan("1", 50, 110000, 10.0)
	err := billingService.MakePayment("1", 110000)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	outstanding, _ := billingService.GetOutstanding("1")
	if outstanding != 5390000 {
		t.Errorf("Expected outstanding amount after payment to be 5390000, got %d", outstanding)
	}
}

func TestMakePaymentInvalidAmount(t *testing.T) {
	store := store.NewInMemoryStore()
	billingService := NewBillingService(store)

	billingService.CreateLoan("1", 50, 110000, 10.0)
	err := billingService.MakePayment("1", 50000)
	if err == nil || err.Error() != "payment amount must be a multiple of the weekly payment" {
		t.Errorf("Expected error 'payment amount must be a multiple of the weekly payment', got %v", err)
	}
}

func TestIsDelinquent(t *testing.T) {
	store := store.NewInMemoryStore()
	billingService := NewBillingService(store)

	billingService.CreateLoan("1", 50, 110000, 10.0)

	// Simulate 1 missed payments
	loan, _ := billingService.GetLoan("1")
	loan.Schedule[0].Paid = true
	loan.Schedule[1].Paid = false

	delinquent, err := billingService.IsDelinquent("1", 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if delinquent {
		t.Errorf("Expected non-delinquent borrower, but got delinquent")
	}

	// Simulate 2 missed payments
	loan, _ = billingService.GetLoan("1")
	loan.Schedule[0].Paid = false
	loan.Schedule[1].Paid = false

	delinquent, _ = billingService.IsDelinquent("1", 3)
	if !delinquent {
		t.Errorf("Expected delinquent borrower, but got non-delinquent")
	}
}
