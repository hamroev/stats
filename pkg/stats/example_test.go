package stats

import (
	"fmt"

	"github.com/hamroev/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   5_00,
			Category: "Cat1",
			Status:   "OK",
		},
		{
			ID:       1,
			Amount:   8_00,
			Category: "Cat1",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   7_00,
			Category: "Cat2",
			Status:   "OK",
		},
		{
			ID:       4,
			Amount:   9_00,
			Category: "Cat2",
			Status:   "FAIL",
		},
		{
			ID:       5,
			Amount:   -3_00,
			Category: "Cat3",
			Status:   "OK",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 425
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   5_00,
			Category: "Cat1",
			Status:   "OK",
		},
		{
			ID:       1,
			Amount:   8_00,
			Category: "Cat1",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   7_00,
			Category: "Cat2",
			Status:   "FAIL",
		},
		{
			ID:       4,
			Amount:   9_00,
			Category: "Cat2",
			Status:   "OK",
		},
		{
			ID:       5,
			Amount:   -3_00,
			Category: "Cat3",
			Status:   "OK",
		},
	}

	fmt.Println(TotalInCategory(payments, "Cat2"))

	//Output: 900
}
