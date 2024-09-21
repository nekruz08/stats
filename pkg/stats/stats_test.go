package stats

import (
	"reflect"
	"testing"

	"github.com/nekruz08/bank/v2/pkg/bank/types"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
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

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}
	// result:=FilterByCategory(payments,"auto")
	// if !reflect.DeepEqual(expected,result){
	// 	t.Error("invalid result")
	// }

	result := FilterByCategory(payments, "auto")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result,expected: %v,actual: %v", expected, result)
	}

}

func TestCategoriesTotal(t *testing.T) {
	payment := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "funn", Amount: 5_000_00},
	}

	expected:=map[types.Category]types.Money{
		"auto":8_000_00,
		"food":2_000_00,
		"funn":5_000_00,
	}

	result:=CategoriesTotal(payment)

	if !reflect.DeepEqual(expected,result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected,result)
	}
}

func TestCategoriesAvg(t *testing.T) {
	payment := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 2, Category: "auto", Amount: 10},
		{ID: 3, Category: "auto", Amount: 40},
	}

	expected:=map[types.Category]types.Money{
		"auto":20,
	}

	result:=CategoriesAvg(payment)

	if !reflect.DeepEqual(expected,result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected,result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto":10,
		"food":20,
		"pood":30,
	}

	second := map[types.Category]types.Money{
		"auto":11,
		"food":22,
		"pood":33,
	}

	expected:=map[types.Category]types.Money{
		"auto":1,
		"food":2,
		"pood":3,
	}

	result:=PeriodsDynamic(first,second)

	if !reflect.DeepEqual(expected,result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected,result)
	}
}