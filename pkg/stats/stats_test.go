package stats

import (
	"reflect"
	"testing"
	"github.com/nekruz08/bank/v2/pkg/types"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len!=0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestFilterByCategory_foundMultple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "food"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "food"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "food"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 4_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 800000,
		"food": 200000,
		"fun":  400000,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 2, Category: "auto", Amount: 20},
		{ID: 3, Category: "food", Amount: 30},
		{ID: 4, Category: "food", Amount: 60},
	}
	expected := map[types.Category]types.Money{
		"auto": 15,
		"food": 45,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first:=map[types.Category]types.Money{
		"auto":10,
		"food":20,
	}
	second:=map[types.Category]types.Money{
		"auto":20,
		"food":20,
	}

	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 0,
	}

	result := PeriodsDynamic(first,second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

	
}
