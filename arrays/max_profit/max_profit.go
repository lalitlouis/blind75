package main

import (
	"fmt"
)

func main() {
	TestMaxProfit()
}

func TestMaxProfit() {
	testCases := []struct {
		name     string
		prices   []int
		expected int
		hasError bool
		desc     string
	}{
		{"basic profit", []int{7, 1, 5, 3, 6, 4}, 5, false, "Buy at 1, sell at 6"},
		{"no profit", []int{7, 6, 4, 3, 2, 1}, 0, false, "Decreasing prices"},
		{"increasing prices", []int{1, 2, 3, 4, 5}, 4, false, "Buy at 1, sell at 5"},
		{"single element", []int{5}, 0, true, "Not enough elements"},
		{"empty array", []int{}, 0, true, "Empty array"},
		{"two elements profit", []int{1, 5}, 4, false, "Simple two element profit"},
		{"two elements no profit", []int{5, 1}, 0, false, "Two elements, no profit"},
		{"same prices", []int{3, 3, 3, 3}, 0, false, "All same prices"},
	}

	for _, tc := range testCases {
		result, err := MaxProfit(tc.prices)

		fmt.Printf("Test: %s - ", tc.desc)

		if tc.hasError {
			if err != nil {
				fmt.Println("PASS (expected error)")
			} else {
				fmt.Println("FAIL (expected error)")
			}
		} else {
			if err != nil {
				fmt.Printf("FAIL (unexpected error: %v)\n", err)
			} else if result == tc.expected {
				fmt.Println("PASS")
			} else {
				fmt.Printf("FAIL (expected %d, got %d)\n", tc.expected, result)
			}
		}
	}
}

func MaxProfit(prices []int) (int, error) {

	if len(prices) < 2 {
		return 0, fmt.Errorf("There should be atleast 2 values in the array")
	}

	min := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < min {
			min = price
		}

		if profit := price - min; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit, nil
}
