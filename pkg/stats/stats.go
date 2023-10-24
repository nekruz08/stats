package stats

import "github.com/nekruz08/bank/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avg:=types.Money(0)
	for _, payment := range payments {
		avg+=payment.Amount
	}
	return avg/types.Money(len(payments))
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment,category types.Category)types.Money  {
	avg:=types.Money(0)
	for _, payment := range payments {
		if payment.Category==category{
			avg+=payment.Amount
		}
	}
	return avg
}

// FilterByCategory возвращает платежи в укащанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category==category{
			filtered=append(filtered, payment)
		}
	}
	return filtered
}

// CategoriesTotal возвращает сумму платежей под каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories:=map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
	}
	return categories
}

// CategoriesAvg cчитает среднюю сумму платежа по каждой категории.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories:=map[types.Category]types.Money{}
	counter:=map[types.Category]types.Money{}
					//мой

	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
		   counter[payment.Category]++
	}

	for key := range categories {
		categories[key]/=counter[key]
	}
	return categories
}
