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
