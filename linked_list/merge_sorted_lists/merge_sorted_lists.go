package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	TestMergeSortedList()
}

func TestMergeSortedList() {
	testCases := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
		desc     string
	}{
		{"example case", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}, "Standard merge"},
		{"empty lists", []int{}, []int{}, []int{}, "Both lists empty"},
		{"one empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}, "First list empty"},
		{"second empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}, "Second list empty"},
		{"different lengths", []int{1, 3, 5}, []int{2, 4, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}, "Different lengths"},
		{"all same values", []int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, "Duplicate values"},
		{"no overlap", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, "No overlapping values"},
		{"single nodes", []int{1}, []int{2}, []int{1, 2}, "Single node in each"},
	}

	for _, tc := range testCases {
		list1 := createList(tc.list1)
		list2 := createList(tc.list2)
		merged := MergeSortedList(list1, list2)
		result := listToArray(merged)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("List1: %s\n", arrayToString(tc.list1))
		fmt.Printf("List2: %s\n", arrayToString(tc.list2))
		fmt.Printf("Expected: %v, Got: %v\n", tc.expected, result)

		if slicesEqual(result, tc.expected) {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

type Node struct {
	Val  int
	Next *Node
}

// Corrected function
func MergeSortedList(list1, list2 *Node) *Node {
	dummy := &Node{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val { // Fixed: <= instead of >=
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Simplified: just attach remaining list
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

// Helper functions
func createList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{Val: values[0]}
	curr := head
	for i := 1; i < len(values); i++ {
		curr.Next = &Node{Val: values[i]}
		curr = curr.Next
	}
	return head
}

func listToArray(head *Node) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func arrayToString(arr []int) string {
	if len(arr) == 0 {
		return "nil"
	}
	strs := make([]string, len(arr))
	for i, v := range arr {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, " -> ") + " -> nil"
}

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
