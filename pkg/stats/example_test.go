package stats

import (
	"fmt"

	"github.com/nekruz08/bank/v2/pkg/bank/types"
)

// Avg testing section
func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 2,
			Status: types.StatusFail,
		},
		{
			Amount: 4,
		},
		{
			Amount: 2,
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
			Status: types.StatusFail,
		},
		{
			Amount: 4,
			Category: "food",
		},
		{
			Amount: 2,
		},
		{
			Amount: 0,
		},
	},"food"))
	//Output:
	// 6

}