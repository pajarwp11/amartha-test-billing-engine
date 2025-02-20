package main

import "errors"

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

	if paymentsMade < numPayments {
		return errors.New("payment exceeds pending dues")
	}

	return nil
}
