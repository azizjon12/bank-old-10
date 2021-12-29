package stats

import "github.com/azizjon12/bank/pkg/bank/types"

// Avg calculates average sum of the payment
func Avg(payments []types.Payment) (amount types.Money) {
	for _, payment := range payments {
		amount += payment.Amount
	}
	return amount / types.Money(len(payments))
}

// TotalInCategory sums up the payments made on specific category
func TotalInCategory(payments []types.Payment, category types.Category) (amount types.Money) {
	for _, payment := range payments {
		if payment.Category == category {
			amount += payment.Amount
		}
	}
	return
}
