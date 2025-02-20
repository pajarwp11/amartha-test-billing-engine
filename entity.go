package main

const (
	DelinquentWeeks = 2
)

type PaymentSchedule struct {
	WeekNumber int
	Amount     int
	Paid       bool
}

type Loan struct {
	ID            string
	TotalWeeks    int
	Outstanding   int
	WeeklyPayment int
	Interest      float64
	Schedule      []PaymentSchedule
}

type LoanStore struct {
	Loans map[string]*Loan
}
