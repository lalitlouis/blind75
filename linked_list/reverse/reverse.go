package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	TestReverseLinkedList()
}

func TestReverseLinkedList() {
	fmt.Println("=== ITERATIVE REVERSE TESTS ===")

	testCases := []struct {
		name     string
		list     []int
		expected []int
		desc     string
	}{
		{"single node", []int{1}, []int{1}, "Single node remains unchanged"},
		{"two nodes", []int{1, 2}, []int{2, 1}, "Simple two node reversal"},
		{"multiple nodes", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, "Standard reversal"},
		{"empty list", []int{}, []int{}, "Empty list test"},
	}

	for _, tc := range testCases {
		head := createList(tc.list)
		reversed, err := ReverseLinkedList(head)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Input: %s\n", listToString(head))

		if len(tc.list) == 0 {
			if err != nil {
				fmt.Println("✅ Expected error for empty list: PASS")
			} else {
				fmt.Println("❌ Expected error for empty list: FAIL")
			}
		} else {
			if err != nil {
				fmt.Printf("Unexpected error: %v\n", err)
			} else {
				result := listToArray(reversed)
				fmt.Printf("Expected: %v, Got: %v\n", tc.expected, result)
				if slicesEqual(result, tc.expected) {
					fmt.Println("✅ PASS")
				} else {
					fmt.Println("❌ FAIL")
				}
			}
		}
	}

	fmt.Println("=== RECURSIVE REVERSE TESTS ===")
	for _, tc := range testCases {
		if len(tc.list) == 0 {
			continue // Skip empty test for recursive
		}

		head := createList(tc.list)
		reversed := RecursiveLinkedList(head)
		result := listToArray(reversed)

		fmt.Printf("Test: %s (Recursive)\n", tc.name)
		fmt.Printf("Expected: %v, Got: %v\n", tc.expected, result)
		if slicesEqual(result, tc.expected) {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
	}
}

type Node struct {
	Val  int
	Next *Node
}

// Corrected iterative function
func ReverseLinkedList(head *Node) (*Node, error) {
	if head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	var prev *Node // Start with nil, not a new node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev // Point to previous, not nil
		prev = curr
		curr = next
	}

	return prev, nil
}

func RecursiveLinkedList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := RecursiveLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
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

func listToString(head *Node) string {
	if head == nil {
		return "nil"
	}

	values := []string{}
	for head != nil {
		values = append(values, strconv.Itoa(head.Val))
		head = head.Next
	}
	return strings.Join(values, " -> ") + " -> nil"
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
