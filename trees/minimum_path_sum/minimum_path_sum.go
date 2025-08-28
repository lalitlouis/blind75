package main

import "fmt"

func main() {
	TestMinimumPathSum()
}

func TestMinimumPathSum() {
	testCases := []struct {
		name     string
		grid     [][]int
		expected int
		desc     string
	}{
		{
			"example 1",
			[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			7,
			"Path: 1→3→1→1→1 = 7",
		},
		{
			"example 2",
			[][]int{{1, 2, 3}, {4, 5, 6}},
			12,
			"Path: 1→2→3→6 = 12",
		},
		{
			"single cell",
			[][]int{{5}},
			5,
			"Only one cell, sum = 5",
		},
		{
			"single row",
			[][]int{{1, 3, 1, 5}},
			10,
			"Single row: 1+3+1+5 = 10",
		},
		{
			"single column",
			[][]int{{1}, {2}, {3}},
			6,
			"Single column: 1+2+3 = 6",
		},
		{
			"all ones",
			[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			5,
			"3x3 grid of ones, minimum path = 5",
		},
		{
			"increasing values",
			[][]int{{1, 4}, {2, 3}},
			6,
			"Path: 1→2→3 = 6",
		},
		{
			"large values",
			[][]int{{10, 20}, {30, 40}},
			60,
			"Path: 10→20→40 = 70 or 10→30→40 = 80, minimum = 70",
		},
	}

	for _, tc := range testCases {
		result := MinimumPathSumHelper(tc.grid)

		fmt.Printf("Test: %s\n", tc.desc)
		fmt.Printf("Grid: %v\n", tc.grid)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)

		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func MinimumPathSumHelper(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	memo := make(map[string]int, 0)
	return MinimumPathSum(grid, rows-1, cols-1, memo)
}

func MinimumPathSum(grid [][]int, i, j int, memo map[string]int) int {

	if i < 0 || j < 0 {
		return 99999999
	}

	if i == 0 && j == 0 {
		return grid[0][0]
	}

	key := fmt.Sprintf("%d,%d", i, j)
	if value, ok := memo[key]; ok {
		return value
	}

	memo[key] = grid[i][j] + min(MinimumPathSum(grid, i-1, j, memo), MinimumPathSum(grid, i, j-1, memo))
	return memo[key]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
