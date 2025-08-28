package main

import "fmt"

func main() {
	TestUniquePaths()
}

func TestUniquePaths() {
	testCases := []struct {
		name     string
		m        int
		n        int
		expected int
		desc     string
	}{
		{"example 1", 3, 7, 28, "3x7 grid has 28 unique paths"},
		{"example 2", 3, 2, 3, "3x2 grid has 3 paths"},
		{"single row", 1, 5, 1, "1x5 grid - only one path (right only)"},
		{"single column", 4, 1, 1, "4x1 grid - only one path (down only)"},
		{"square small", 2, 2, 2, "2x2 grid - 2 paths: right-down or down-right"},
		{"square medium", 3, 3, 6, "3x3 grid has 6 unique paths"},
		{"minimum case", 1, 1, 1, "1x1 grid - one path (stay put)"},
		{"rectangular", 4, 3, 10, "4x3 grid has 10 unique paths"},
		{"larger case", 5, 5, 70, "5x5 grid has 70 unique paths"},
		{"edge case", 7, 3, 28, "7x3 grid has 28 unique paths"},
	}

	for _, tc := range testCases {
		result := UniquePathsHelper(tc.m, tc.n)

		fmt.Printf("Test: %s\n", tc.desc)
		fmt.Printf("Input: m=%d, n=%d\n", tc.m, tc.n)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)

		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func UniquePathsHelper(m, n int) int {
	memo := make(map[string]int)
	return uniquePaths(m-1, n-1, memo)
}

func uniquePaths(m, n int, memo map[string]int) int {

	if m < 0 || n < 0 {
		return 0
	}

	if m == 0 && n == 0 {
		return 1
	}

	key := fmt.Sprintf("%d,%d", m, n)
	if value, ok := memo[key]; ok {
		return value
	}

	memo[key] = uniquePaths(m-1, n, memo) + uniquePaths(m, n-1, memo)
	return memo[key]
}
