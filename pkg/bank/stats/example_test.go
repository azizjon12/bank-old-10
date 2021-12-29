package stats

import (
	"fmt"

	"github.com/azizjon12/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       14,
			Amount:   5_00,
			Category: "shop",
		},
		{
			ID:       21,
			Amount:   200_00,
			Category: "restaurant",
		},
		{
			ID:       22,
			Amount:   20_00,
			Category: "cafe",
		},
	}

	fmt.Println(Avg(payments))
	// Output: 7500
}
