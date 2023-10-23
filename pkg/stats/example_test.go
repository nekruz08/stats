package stats

import (
	"fmt"

	"github.com/nekruz08/bank/v2/pkg/types")

func ExampleAvg() {
	payments:=[]types.Payment{
		{
			ID: 1,
			Amount: 2,
			Category: "restraunt",
		},
		{
			ID: 2,
			Amount: 2,
			Category: "food",
		},
		{
			ID: 2,
			Amount: 2,
			Category: "transport",
		},
	}
	fmt.Println(Avg(payments))
	//Output:2
}

func ExampleTotalInCategory() {
	payments:=[]types.Payment{
		{
			ID: 1,
			Amount: 2,
			Category: "restraunt",
		},
		{
			ID: 2,
			Amount: 2,
			Category: "food",
		},
		{
			ID: 3,
			Amount: 2,
			Category: "transport",
		},
		{
			ID: 4,
			Amount: 2,
			Category: "transport",
		},
		{
			ID: 5,
			Amount: 2,
			Category: "transport",
		},
	}
	fmt.Println(TotalInCategory(payments,"food"))
	//Output:2
}