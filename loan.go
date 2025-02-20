package main

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
