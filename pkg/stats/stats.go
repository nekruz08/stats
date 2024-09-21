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
		if payment.Status == types.StatusFail {
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
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}

// FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment,category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category==category{
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

// CategoriesTotal возвращает сумму платежей по каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories:=map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
	}
	return categories
}

// CategoriesAvg считает среднюю сумму платежа по каждой категории:
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	counter := map[types.Category]int{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		counter[payment.Category]++
	}

	// Подсчитываем среднее значение
	for category := range categories {
		categories[category] /= types.Money(counter[category])
	}

	return categories
}

// PeriodsDynamic сравнивает расходы по категориям за два периода
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money,) map[types.Category]types.Money {
	for k, v := range first {
		second[k]-=v
	}
	return second
}