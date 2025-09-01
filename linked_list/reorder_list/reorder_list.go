package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	TestReorderList()
}

func TestReorderList() {
	testCases := []struct {
		name     string
		input    []int
		expected []int
		desc     string
	}{
		{"odd length example", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}, "Standard 5-node reorder"},
		{"even length", []int{1, 2, 3, 4}, []int{1, 4, 2, 3}, "4-node reorder"},
		{"two nodes", []int{1, 2}, []int{1, 2}, "Two nodes - no change needed"},
		{"single node", []int{1}, []int{1}, "Single node unchanged"},
		{"three nodes", []int{1, 2, 3}, []int{1, 3, 2}, "Three nodes"},
		{"six nodes", []int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}, "Even length - 6 nodes"},
		{"seven nodes", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 7, 2, 6, 3, 5, 4}, "Odd length - 7 nodes"},
	}

	for _, tc := range testCases {
		head := createList(tc.input)
		result := ReorderList(head)
		resultArray := listToArray(result)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Input: %s\n", arrayToString(tc.input))
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

func arrayToString(arr []int) string {
	if len(arr) == 0 {
		return "[]"
	}
	strs := make([]string, len(arr))
	for i, v := range arr {
		strs[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(strs, "->") + "]"
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

func ReorderList(list *Node) *Node {
	middleNode := FindMiddle(list)
	secondHalf := middleNode.Next
	middleNode.Next = nil
	reversedList := Reverse(secondHalf)
	return MergeLists(list, reversedList)
}

func FindMiddle(node *Node) *Node {

	slow := node
	fast := node.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func Reverse(node *Node) *Node {

	var prev *Node
	curr := node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func MergeLists(list1, list2 *Node) *Node {

	dummy := &Node{}
	list := dummy

	for list1 != nil && list2 != nil {
		list.Next = list1
		list1 = list1.Next
		list = list.Next
		list.Next = list2
		list2 = list2.Next
		list = list.Next
	}

	if list1 != nil {
		list.Next = list1
	} else {
		list.Next = list2
	}

	return dummy.Next
}
