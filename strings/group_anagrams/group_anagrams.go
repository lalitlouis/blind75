package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	TestGroupAnagrams()
}

func TestGroupAnagrams() {
	testCases := []struct {
		name     string
		strs     []string
		expected int // We'll check count of groups instead of exact match
		desc     string
	}{
		{"basic example", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, 3, "Three groups: [eat,tea,ate], [tan,nat], [bat]"},
		{"empty string", []string{""}, 1, "Single empty string"},
		{"single char", []string{"a"}, 1, "Single character"},
		{"no anagrams", []string{"abc", "def", "ghi"}, 3, "All different, no anagrams"},
		{"all anagrams", []string{"abc", "bca", "cab"}, 1, "All are anagrams of each other"},
		{"empty array", []string{}, 0, "Empty input"},
		{"duplicates", []string{"ab", "ba", "ab"}, 1, "Duplicates should be in same group"},
	}

	for _, tc := range testCases {
		result := GroupAnagrams(tc.strs)

		fmt.Printf("Test: %s\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.strs)
		fmt.Printf("Output groups: %d\n", len(result))
		fmt.Printf("Expected groups: %d\n", tc.expected)

		if len(result) == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Printf("FAIL - got %d groups, expected %d\n\n", len(result), tc.expected)
		}
	}
}

func GroupAnagrams(strs []string) [][]string {

	strMap := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)

		if value, ok := strMap[sortedStr]; ok {
			value = append(value, str)
			strMap[sortedStr] = value
		} else {
			strMap[sortedStr] = []string{str}
		}
	}

	groupAnagrams := make([][]string, 0)

	for _, v := range strMap {
		groupAnagrams = append(groupAnagrams, v)
	}

	return groupAnagrams
}

func sortString(str string) string {

	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
