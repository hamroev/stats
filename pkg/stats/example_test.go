package stats

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   5_00,
			Category: "Cat1",
		},
		{
			ID:       1,
			Amount:   8_00,
			Category: "Cat1",
		},
		{
			ID:       3,
			Amount:   7_00,
			Category: "Cat2",
		},
		{
			ID:       4,
			Amount:   9_00,
			Category: "Cat2",
		},
		{
			ID:       5,
			Amount:   -3_00,
			Category: "Cat3",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 520
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   5_00,
			Category: "Cat1",
		},
		{
			ID:       1,
			Amount:   8_00,
			Category: "Cat1",
		},
		{
			ID:       3,
			Amount:   7_00,
			Category: "Cat2",
		},
		{
			ID:       4,
			Amount:   9_00,
			Category: "Cat2",
		},
		{
			ID:       5,
			Amount:   -3_00,
			Category: "Cat3",
		},
	}

	fmt.Println(TotalInCategory(payments, "Cat2"))

	//Output: 1600
}
