package stats

import "github.com/nekruz08/bank/v2/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	count := types.Money(0)
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Amount <= 0 {
			continue
		}
		if payment.Status==types.StatusFail {
			continue
		}
		count++
		sum += payment.Amount
	}
	return sum / count
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status==types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
