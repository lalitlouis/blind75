package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	TestRemoveNthNode()
}

func TestRemoveNthNode() {
	testCases := []struct {
		name     string
		list     []int
		k        int
		expected []int
		desc     string
	}{
		{"remove 2nd from end", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}, "Remove node 4"},
		{"remove head", []int{1, 2, 3}, 3, []int{2, 3}, "Remove first node"},
		{"remove tail", []int{1, 2, 3}, 1, []int{1, 2}, "Remove last node"},
		{"single node", []int{1}, 1, []int{}, "Remove only node"},
		{"two nodes remove first", []int{1, 2}, 2, []int{2}, "Remove head of two nodes"},
		{"two nodes remove second", []int{1, 2}, 1, []int{1}, "Remove tail of two nodes"},
		{"middle node", []int{1, 2, 3, 4, 5, 6}, 3, []int{1, 2, 3, 5, 6}, "Remove middle node"},
	}

	for _, tc := range testCases {
		head := createList(tc.list)
		result, _ := RemoveNthNode(head, tc.k)
		resultArray := listToArray(result)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Input: %s, k=%d\n", arrayToString(tc.list), tc.k)
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
	return "[" + strings.Join(strs, ",") + "]"
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

func RemoveNthNode(node *Node, k int) (*Node, error) {

	if node == nil {
		return nil, errors.New("linked list is empty")
	}

	dummy := &Node{Next: node}
	slow := dummy
	fast := dummy

	for i := 0; i <= k; i++ {
		if fast == nil {
			return nil, errors.New("fewer than k elements")
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next, nil

}
