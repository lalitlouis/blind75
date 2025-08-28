package main

import "fmt"

func main() {
	TestHouseRobber()
}

func TestHouseRobber() {
	testCases := []struct {
		name     string
		nums     []int
		expected int
		desc     string
	}{
		{"example 1", []int{1, 2, 3, 1}, 4, "Rob houses 0 and 2: 1 + 3 = 4"},
		{"example 2", []int{2, 7, 9, 3, 1}, 12, "Rob houses 0, 2, 4: 2 + 9 + 1 = 12"},
		{"single house", []int{5}, 5, "Only one house, rob it"},
		{"two houses", []int{2, 3}, 3, "Rob the better of two houses"},
		{"two equal houses", []int{5, 5}, 5, "Rob either house when equal"},
		{"three houses", []int{2, 1, 1}, 2, "Rob only the first house"},
		{"alternating high low", []int{5, 1, 3, 1, 2}, 8, "Rob houses 0, 2, 4: 5 + 3 + 2 = 8"},
		{"increasing values", []int{1, 2, 3, 4, 5}, 9, "Rob houses 1, 3: 2 + 4 = 6... no wait 0,2,4: 1+3+5=9"},
		{"decreasing values", []int{5, 4, 3, 2, 1}, 9, "Rob houses 0, 2, 4: 5 + 3 + 1 = 9"},
		{"empty array", []int{}, 0, "No houses to rob"},
		{"all same values", []int{3, 3, 3, 3, 3}, 9, "Rob houses 0, 2, 4: 3 + 3 + 3 = 9"},
		{"large gap", []int{100, 1, 1, 100}, 200, "Rob houses 0 and 3: 100 + 100 = 200"},
	}

	for _, tc := range testCases {
		result := HouseRobber(tc.nums)

		fmt.Printf("Test: %s\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.nums)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)

		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Printf("FAIL\n\n")
		}
	}
}

func HouseRobber(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	memo := make(map[int]int, 0)

	return robHelper(nums, len(nums)-1, memo)

}

func robHelper(nums []int, i int, memo map[int]int) int {
	if i < 0 {
		return 0
	} else if i == 0 {
		return nums[0]
	}

	if value, ok := memo[i]; ok {
		return value
	}

	memo[i] = max(nums[i]+robHelper(nums, i-2, memo), robHelper(nums, i-1, memo))

	return memo[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
