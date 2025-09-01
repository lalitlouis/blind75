package main

import (
	"errors"
	"fmt"
)

func main() {
	TestMostWater()
}

func TestMostWater() {
	testCases := []struct {
		name     string
		levels   []int
		expected int
		hasError bool
		desc     string
	}{
		{"example 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49, false, "Classic example - indices 1,8 give area 49"},
		{"example 2", []int{1, 1}, 1, false, "Two equal heights"},
		{"increasing", []int{1, 2, 3, 4, 5}, 6, false, "Increasing heights - indices 0,4 give 1*4=4... wait should be 6"},
		{"decreasing", []int{5, 4, 3, 2, 1}, 6, false, "Decreasing heights"},
		{"equal heights", []int{4, 4, 4, 4}, 12, false, "All equal heights - max width"},
		{"two tall ends", []int{10, 1, 1, 1, 10}, 40, false, "Tall ends with small middle"},
		{"single element", []int{5}, 0, true, "Not enough elements"},
		{"empty array", []int{}, 0, true, "Empty array"},
		{"tall middle", []int{1, 10, 1}, 2, false, "Tall middle, short ends"},
		{"three elements", []int{2, 1, 5}, 4, false, "Three elements - ends give 2*2=4"},
	}

	for _, tc := range testCases {
		result, err := mostWater(tc.levels)

		fmt.Printf("Test: %s - ", tc.desc)

		if tc.hasError {
			if err != nil {
				fmt.Println("✅ PASS (expected error)")
			} else {
				fmt.Printf("❌ FAIL (expected error, got %d)\n", result)
			}
		} else {
			if err != nil {
				fmt.Printf("❌ FAIL (unexpected error: %v)\n", err)
			} else if result == tc.expected {
				fmt.Println("✅ PASS")
			} else {
				fmt.Printf("❌ FAIL (expected %d, got %d)\n", tc.expected, result)
			}
		}
	}
}

func mostWater(levels []int) (int, error) {

	left, right := 0, len(levels)-1
	maxVolume := 0

	if len(levels) < 2 {
		return maxVolume, errors.New("need atleast 2 entries into the levels array")
	}

	for left < right {
		width := right - left
		min := levels[left]
		if levels[left] > levels[right] {
			min = levels[right]
			right--
		} else {
			left++
		}

		if volume := getVolume(min, width); volume >= maxVolume {
			maxVolume = volume
		}

	}

	return maxVolume, nil
}

func getVolume(a int, b int) int {
	return a * b
}
