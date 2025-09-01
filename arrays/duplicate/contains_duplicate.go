package main

import (
	"errors"
	"fmt"
)

func main() {
	TestContainsDuplicate()
}

func TestContainsDuplicate() {
	// Test cases: input, expected
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		// Basic cases
		{"Has duplicate", []int{1, 2, 3, 1}, true},
		{"No duplicate", []int{1, 2, 3, 4}, false},
		{"All same", []int{1, 1, 1}, true},

		// Edge cases
		{"Single element", []int{1}, false},
		{"Empty array", []int{}, false},
		{"Two same", []int{1, 1}, true},
		{"Two different", []int{1, 2}, false},

		// Larger arrays
		{"Large with duplicate", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}, true},
		{"Large no duplicate", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},

		// Negative numbers
		{"Negative duplicate", []int{-1, -2, -3, -1}, true},
		{"Mixed with duplicate", []int{-1, 0, 1, -1}, true},
	}

	// Run tests
	for _, tc := range testCases {
		result, _ := ContainsDuplicate(tc.nums)
		if result == tc.expected {
			fmt.Printf("✅ %s: PASS\n", tc.name)
		} else {
			fmt.Printf("❌ %s: FAIL (got %v, want %v)\n", tc.name, result, tc.expected)
		}
	}
}

func ContainsDuplicate(nums []int) (bool, error) {

	if len(nums) < 2 {
		return false, errors.New("only 2 element")
	}

	numsMap := make(map[int]int)

	for idx, num := range nums {

		if _, ok := numsMap[num]; ok {
			return true, nil
		}
		numsMap[num] = idx
	}

	return false, errors.New("does not contain duplicate")
}
