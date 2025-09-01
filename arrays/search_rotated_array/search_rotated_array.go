package main

import (
	"fmt"
)

func main() {
	TestSearchRotatedArray()
}

func TestSearchRotatedArray() {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
		desc     string
	}{
		{"example 1", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4, "target in right half"},
		{"example 2", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1, "target not found"},
		{"single element found", []int{1}, 1, 0, "single element match"},
		{"single element not found", []int{1}, 0, -1, "single element no match"},
		{"target in left half", []int{4, 5, 6, 7, 0, 1, 2}, 5, 1, "target in sorted left half"},
		{"no rotation", []int{1, 2, 3, 4, 5}, 3, 2, "regular sorted array"},
		{"target at pivot", []int{4, 5, 6, 7, 0, 1, 2}, 7, 3, "target at rotation point"},
		{"two elements", []int{1, 3}, 3, 1, "two element array"},
		{"reversed array", []int{3, 1}, 1, 1, "completely rotated"},
	}

	for _, tc := range testCases {
		result := SearchRotatedArray(tc.nums, tc.target)

		fmt.Printf("Test: %s\n", tc.desc)
		fmt.Printf("Input: nums=%v, target=%d\n", tc.nums, tc.target)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)

		if result == tc.expected {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
	}
}
func SearchRotatedArray(nums []int, target int) int {

	left, mid, right := 0, len(nums)/2, len(nums)-1

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		}
		// left half is sorted
		if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] >= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	if nums[mid] == target {
		return mid
	}
	return -1

}

func getGpuResources(nums []int, target int) bool {

	left, mid, right := 0, len(nums)/2, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		if nums[left] <= nums[mid] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	if nums[mid] == target {
		return true
	}

	return false
}
