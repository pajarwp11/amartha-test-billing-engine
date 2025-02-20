package loan

import (
	"testing"
)

func TestNewLoan(t *testing.T) {
	loan := NewLoan("1", 50, 110000, 10.0)
	if loan.ID != "1" {
		t.Errorf("Expected loan ID to be '1', got %s", loan.ID)
	}
	if loan.TotalWeeks != 50 {
		t.Errorf("Expected total weeks to be 50, got %d", loan.TotalWeeks)
	}
	if loan.WeeklyPayment != 110000 {
		t.Errorf("Expected weekly payment to be 110000, got %d", loan.WeeklyPayment)
	}
	if loan.Outstanding != 5500000 {
		t.Errorf("Expected outstanding amount to be 5500000, got %d", loan.Outstanding)
	}
}

func TestMakePayment(t *testing.T) {
	loan := NewLoan("1", 50, 110000, 10.0)
	beforeOutstanding := loan.GetOutstanding()
	if beforeOutstanding != 5500000 {
		t.Errorf("Expected outstanding amount before payment to be 5500000, got %d", beforeOutstanding)
	}

	err := loan.MakePayment(110000)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	afterOutstanding := loan.GetOutstanding()
	if afterOutstanding != 5390000 {
		t.Errorf("Expected outstanding amount after payment to be 5390000, got %d", afterOutstanding)
	}
}

func TestMakePaymentInvalidAmount(t *testing.T) {
	loan := NewLoan("1", 50, 110000, 10.0)
	err := loan.MakePayment(50000)
	if err == nil || err.Error() != "payment amount must be a multiple of the weekly payment" {
		t.Errorf("Expected error 'payment amount must be a multiple of the weekly payment', got %v", err)
	}
}

func TestIsDelinquent(t *testing.T) {
	loan := NewLoan("1", 50, 110000, 10.0)
	if loan.IsDelinquent(3) {
		t.Errorf("Expected non-delinquent borrower, but got delinquent")
	}
	loan.Schedule[0].Paid = false
	loan.Schedule[1].Paid = false
	if !loan.IsDelinquent(3) {
		t.Errorf("Expected delinquent borrower, but got non-delinquent")
	}
}
