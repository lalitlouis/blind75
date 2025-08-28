package main

import "fmt"

func main() {
	TestLongestSubstringWithoutRepeatingCharacters()
}

func TestLongestSubstringWithoutRepeatingCharacters() {
	testCases := []struct {
		name     string
		str      string
		expected int
		desc     string
	}{
		{"example 1", "abcabcbb", 3, "abc - length 3"},
		{"example 2", "bbbbb", 1, "all same chars - length 1"},
		{"example 3", "pwwkew", 3, "wke - length 3"},
		{"empty string", "", 0, "empty string"},
		{"single char", "a", 1, "single character"},
		{"no repeats", "abcdef", 6, "all unique characters"},
		{"all same", "aaaa", 1, "all same character"},
		{"two chars", "au", 2, "two different chars"},
		{"complex", "dvdf", 3, "vdf - length 3"},
		{"with spaces", "a b c a b c", 3, "with spaces"},
		{"long unique", "abcdefghij", 10, "long unique string"},
	}

	for _, tc := range testCases {
		result := GetLongestSubstring(tc.str)

		fmt.Printf("Test: %s\n", tc.desc)
		fmt.Printf("Input: \"%s\"\n", tc.str)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)

		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Printf("FAIL\n\n")
		}
	}
}

func GetLongestSubstring(str string) int {

	left := 0
	right := 0

	strMap := make(map[byte]int)
	maxLength := 0
	for right < len(str) {
		char := str[right]
		if idx, ok := strMap[char]; ok {
			left = max(left, idx+1)
		}
		strMap[char] = right
		currentLength := right - left + 1
		maxLength = max(currentLength, maxLength)
		right++
	}

	return maxLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
