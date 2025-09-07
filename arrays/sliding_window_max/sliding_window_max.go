package main

import "fmt"

func main() {
	TestSlidingWindowMax()
}

func TestSlidingWindowMax() {
	// Test cases: input array, k, expected output
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		// Basic cases
		{"Classic example", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"Your example", []int{12, 4, 5, 2, 34, 45, 5, 5}, 3, []int{12, 5, 34, 45, 45, 45}},
		{"Increasing sequence", []int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
		{"Decreasing sequence", []int{5, 4, 3, 2, 1}, 3, []int{5, 4, 3}},

		// Edge cases
		{"Single element", []int{5}, 1, []int{5}},
		{"Two elements k=1", []int{1, 3}, 1, []int{1, 3}},
		{"Two elements k=2", []int{1, 3}, 2, []int{3}},
		{"All same elements", []int{2, 2, 2, 2}, 2, []int{2, 2, 2}},

		// Different window sizes
		{"Window size 1", []int{4, 3, 2, 1}, 1, []int{4, 3, 2, 1}},
		{"Window size equals array", []int{1, 3, 2}, 3, []int{3}},
		{"Large window", []int{1, 3, -1, -3, 5, 3, 6, 7}, 4, []int{3, 5, 5, 6, 7}},

		// Special cases
		{"With negatives", []int{-1, -2, -3, -4}, 2, []int{-1, -2, -3}},
		{"Mixed positive negative", []int{1, -1, 0, 3, -2}, 3, []int{1, 3, 3}},

		// Array smaller than k
		{"Array smaller than k", []int{1, 2}, 3, []int{}},
	}

	// Run tests
	for _, tc := range testCases {
		result := SlidingWindowMax(tc.nums, tc.k)
		if slicesEqual(result, tc.expected) {
			fmt.Printf("✅ %s: PASS (result: %v)\n", tc.name, result)
		} else {
			fmt.Printf("❌ %s: FAIL (got %v, want %v)\n", tc.name, result, tc.expected)
		}
	}
}

// nums: list of numbers [12,4,5,2,34,45,5,5]
// k: chunk size or window size 3
// expected output: 12,5,34,45,45,45
func SlidingWindowMaxKNSolution(nums []int, k int) []int {
	result := []int{}
	if len(nums) < k {
		return []int{}
	}

	for i := 0; i <= len(nums)-k; i++ { // Fixed: should be <= not <
		result = append(result, max(nums[i:i+k]))
	}

	return result
}

func max(nums []int) int {
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

// Helper function to compare slices
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// nums = array of ints [1,2,3,4,5]
// k = window size 3
// [3,4,5]
func SlidingWindowMax(nums []int, k int) []int {

	result := []int{}
	if len(nums) < k {
		return result
	}
	prevIndices := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(prevIndices) > 0 && prevIndices[0] <= i-k {
			prevIndices = prevIndices[1:]
		}

		for len(prevIndices) > 0 && nums[prevIndices[len(prevIndices)-1]] < nums[i] {
			prevIndices = prevIndices[:len(prevIndices)-1]
		}

		prevIndices = append(prevIndices, i)

		if i >= k-1 {
			result = append(result, nums[prevIndices[0]])
		}
	}

	return result

}
