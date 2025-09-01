package main

import "fmt"

func main() {
	TestCycleDetection()
}

func TestCycleDetection() {
	testCases := []struct {
		name     string
		setup    func() *Node
		expected bool
		desc     string
	}{
		{"no cycle", createLinearList, false, "Linear list 1->2->3->4"},
		{"cycle at end", createCycleAtEnd, true, "Last node points back to middle"},
		{"cycle at head", createCycleAtHead, true, "Last node points back to head"},
		{"self loop", createSelfLoop, true, "Node points to itself"},
		{"empty list", createEmpty, false, "Null head"},
		{"single node", createSingleNode, false, "Single node with no cycle"},
		{"two node cycle", createTwoNodeCycle, true, "Two nodes pointing to each other"},
		{"large cycle", createLargeCycle, true, "Cycle in larger list"},
	}

	for _, tc := range testCases {
		head := tc.setup()
		result := CycleDetection(head)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %t, Got: %t\n", tc.expected, result)

		if result == tc.expected {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
	}
}

// Test setup functions
func createLinearList() *Node {
	// 1 -> 2 -> 3 -> 4 -> nil
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = &Node{Val: 4}
	return head
}

func createCycleAtEnd() *Node {
	// 3 -> 2 -> 0 -> -4
	//      ^         |
	//      |_________|
	head := &Node{Val: 3}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 0}
	node4 := &Node{Val: -4}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Cycle back to node2

	return head
}

func createCycleAtHead() *Node {
	// 1 -> 2 -> 3
	// ^         |
	// |_________|
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = head // Cycle back to head

	return head
}

func createSelfLoop() *Node {
	// 1 -> (points to itself)
	head := &Node{Val: 1}
	head.Next = head
	return head
}

func createEmpty() *Node {
	return nil
}

func createSingleNode() *Node {
	return &Node{Val: 1}
}

func createTwoNodeCycle() *Node {
	// 1 <-> 2 (cycle between two nodes)
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node1.Next = node2
	node2.Next = node1
	return node1
}

func createLargeCycle() *Node {
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6
	//           ^              |
	//           |______________|
	nodes := make([]*Node, 6)
	for i := 0; i < 6; i++ {
		nodes[i] = &Node{Val: i + 1}
	}

	for i := 0; i < 5; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[5].Next = nodes[2] // Cycle back to node 3

	return nodes[0]
}

type Node struct {
	Val  int
	Next *Node
}

func CycleDetection(node *Node) bool {

	if node == nil || node.Next == nil {
		return false
	}

	slow := node
	fast := node.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next

	}
	return false
}
