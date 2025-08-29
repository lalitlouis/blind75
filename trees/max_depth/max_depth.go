package main

import "fmt"

func main() {
	TestMaxDepth()
}

func TestMaxDepth() {
	testCases := []struct {
		name     string
		tree     *TreeNode
		expected int
		desc     string
	}{
		{"empty tree", nil, 0, "Nil tree has depth 0"},
		{"single node", NewNode(1), 1, "Single node has depth 1"},
		{"example tree", buildExampleTree(), 3, "Tree from your example"},
		{"left skewed", buildLeftSkewed(), 3, "All nodes go left"},
		{"right skewed", buildRightSkewed(), 3, "All nodes go right"},
		{"balanced", buildBalanced(), 3, "Perfectly balanced tree"},
		{"uneven", buildUneven(), 4, "Unbalanced tree"},
	}

	for _, tc := range testCases {
		result := MaxDepth(tc.tree)
		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)
		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func buildExampleTree() *TreeNode {
	//     3
	//    / \
	//   9  20
	//     /  \
	//    15   7
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	return root
}

func buildLeftSkewed() *TreeNode {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Left.Left = NewNode(3)
	return root
}

func buildRightSkewed() *TreeNode {
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Right = NewNode(3)
	return root
}

func buildBalanced() *TreeNode {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	return root
}

func buildUneven() *TreeNode {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Left.Left = NewNode(3)
	root.Left.Left.Left = NewNode(4)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}

}

func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
