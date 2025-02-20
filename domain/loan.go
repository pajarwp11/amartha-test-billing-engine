package loan

import "errors"

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

func NewLoan(id string, totalWeeks, weeklyPayment int, interest float64) *Loan {
	schedule := make([]PaymentSchedule, totalWeeks)
	for i := 0; i < totalWeeks; i++ {
		schedule[i] = PaymentSchedule{
			WeekNumber: i + 1,
			Amount:     weeklyPayment,
			Paid:       false,
		}
	}

	return &Loan{
		ID:            id,
		TotalWeeks:    totalWeeks,
		Outstanding:   totalWeeks * weeklyPayment,
		WeeklyPayment: weeklyPayment,
		Interest:      interest,
		Schedule:      schedule,
	}
}

func (l *Loan) GetOutstanding() int {
	return l.Outstanding
}

func (l *Loan) IsDelinquent(currentWeek int) bool {
	missedCount := 0
	for i := currentWeek - 1; i >= 0 && i >= currentWeek-DelinquentWeeks; i-- {
		if !l.Schedule[i].Paid {
			missedCount++
		}
	}
	return missedCount >= DelinquentWeeks
}

func (l *Loan) MakePayment(amount int) error {
	if amount%l.WeeklyPayment != 0 {
		return errors.New("payment amount must be a multiple of the weekly payment")
	}

	numPayments := amount / l.WeeklyPayment
	paymentsMade := 0

	for i := 0; i < len(l.Schedule) && paymentsMade < numPayments; i++ {
		if !l.Schedule[i].Paid {
			l.Schedule[i].Paid = true
			l.Outstanding -= l.WeeklyPayment
			paymentsMade++
		}
	}

	return nil
}
