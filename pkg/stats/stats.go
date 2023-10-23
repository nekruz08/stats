package stats

import "github.com/nekruz08/bank/v2/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avg:=types.Money(0)
	for _, payment := range payments {
		if payment.Status!=types.StatusFail{
			avg+=payment.Amount
		}
	}
	return avg/types.Money(len(payments))
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment,category types.Category)types.Money  {
	avg:=types.Money(0)
	for _, payment := range payments {
		if payment.Category==category&&payment.Status!=types.StatusFail{
			avg+=payment.Amount
		}
	}
	return avg
}