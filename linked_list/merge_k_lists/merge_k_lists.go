package main

import (
	"fmt"
)

func main() {
	TestMergeKLists()
}

func TestMergeKLists() {
	testCases := []struct {
		name     string
		lists    [][]int
		expected []int
		desc     string
	}{
		{"example", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}, "Standard k lists merge"},
		{"empty lists", [][]int{}, []int{}, "No lists provided"},
		{"single list", [][]int{{1, 2, 3}}, []int{1, 2, 3}, "Only one list"},
		{"empty and non-empty", [][]int{{}, {1, 2}}, []int{1, 2}, "Mix of empty and non-empty"},
		{"all empty", [][]int{{}, {}, {}}, []int{}, "All lists empty"},
		{"two lists", [][]int{{1, 3, 5}, {2, 4, 6}}, []int{1, 2, 3, 4, 5, 6}, "Two sorted lists"},
		{"odd number", [][]int{{1}, {2}, {3}}, []int{1, 2, 3}, "Three single-element lists"},
		{"duplicates", [][]int{{1, 1, 2}, {1, 1, 3}, {2, 2}}, []int{1, 1, 1, 1, 2, 2, 2, 3}, "Lists with duplicates"},
		{"single elements", [][]int{{5}, {1}, {3}}, []int{1, 3, 5}, "Multiple single elements"},
	}

	for _, tc := range testCases {
		// Convert arrays to linked lists
		lists := make([]*Node, len(tc.lists))
		for i, arr := range tc.lists {
			lists[i] = createList(arr)
		}

		result := mergeKLists(lists)
		resultArray := listToArray(result)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Input: %v\n", tc.lists)
		fmt.Printf("Expected: %v, Got: %v\n", tc.expected, resultArray)

		if slicesEqual(resultArray, tc.expected) {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
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

type Node struct {
	Val  int
	Next *Node
}

// Your mergeKLists and merge2Lists functions...

func mergeKLists(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedLists := make([]*Node, 0)

		for i := 0; i < len(lists); i += 2 {
			list1 := lists[i]
			var list2 *Node
			if i+1 < len(lists) {
				list2 = lists[i+1]
			}
			mergedList := merge2Lists(list1, list2)
			mergedLists = append(mergedLists, mergedList)
		}
		lists = mergedLists
	}

	return lists[0]
}

func merge2Lists(list1, list2 *Node) *Node {

	dummy := &Node{}
	list := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list.Next = list1
			list1 = list1.Next
		} else {
			list.Next = list2
			list2 = list2.Next
		}
		list = list.Next
	}

	if list1 != nil {
		list.Next = list1
	}

	if list2 != nil {
		list.Next = list2
	}

	return dummy.Next
}
