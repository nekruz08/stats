package stats

import (
	"fmt"

	"github.com/nekruz08/bank/pkg/bank/types"
)

// Avg testing section
func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 2,
		},
		{
			Amount: 4,
		},
		{
			Amount: 0,
		},
		{
			Amount: 0,
		},
	}))
	//Output:
	// 3
}

// TotalInCategory testin section
func ExampleTotalInCategory() {
	fmt.Println(TotalInCategory([]types.Payment{
		{
			Amount: 2,
			Category: "food",
		},
		{
			Amount: 4,
			Category: "food",
		},
		{
			Amount: 0,
		},
		{
			Amount: 0,
		},
	},"food"))
	//Output:
	// 6

}