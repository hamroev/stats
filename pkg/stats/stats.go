package stats

import (
	"github.com/hamroev/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avg := types.Money(0)
	sum := types.Money(0)
	count := int(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		count++
		sum += payment.Amount
	}
	avg = types.Money(int(sum) / count)

	return avg
}

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}
		sum += payment.Amount
	}

	return sum
}

// CategoriesTotal возвращает сумму платежей по каждой категории
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}
