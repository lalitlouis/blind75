package main

import "fmt"

func TestTwoSum() {

	fmt.Println("Finding the indices of the sum in the current array")

	testCases := []struct {
		nums     []int
		target   int
		expected []int
		desc     string
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}, "Basic case"},
		{[]int{3, 2, 4}, 6, []int{1, 2}, "Non-adjacent elements"},
		{[]int{3, 3}, 6, []int{0, 1}, "Duplicate elements"},
		{[]int{1, 2, 3}, 7, nil, "No solution exists"},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}, "Negative numbers"},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}, "Zeros"},
	}

	for i, tc := range testCases {
		result, found := TwoSum(tc.nums, tc.target)

		fmt.Printf("\nTest %d: %s\n", i+1, tc.desc)
		fmt.Printf("Input: nums=%v, target=%d\n", tc.nums, tc.target)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Got: %v (found: %t)\n", result, found)

		if found != nil && tc.expected == nil {
			fmt.Println("PASS")
		} else if found == nil && tc.expected != nil && len(result) == 2 {
			if tc.nums[result[0]]+tc.nums[result[1]] == tc.target {
				fmt.Println("PASS")
			} else {
				fmt.Println("FAIL")
			}
		} else {
			fmt.Println("FAIL")
		}
	}

}

func TwoSum(nums []int, target int) ([]int, error) {

	numsMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if value, ok := numsMap[diff]; ok {
			fmt.Println("Found indices")
			return []int{value, i}, nil
		}
		numsMap[num] = i
	}

	return []int{}, fmt.Errorf("None of the indices add up to the target number %d", target)
}
