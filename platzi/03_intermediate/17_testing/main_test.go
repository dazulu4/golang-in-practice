package main

import "testing"

// Individual Test
// func TestSum(t *testing.T) {
// 	total := Sum(5, 5)
// 	if total != 10 {
// 		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
// 	}
// }

// Table Test Cases
func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 5, 5},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d, expected %d", max, item.n)
		}
	}
}

//go test -coverprofile=coverage.out
//go tool cover -func=coverage.out
//go tool cover -html=coverage.out
