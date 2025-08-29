package main

import "fmt"

func main() {
	TestSameTrees()
}

func TestSameTrees() {
	testCases := []struct {
		name     string
		tree1    *TreeNode
		tree2    *TreeNode
		expected bool
		desc     string
	}{
		{"both nil", nil, nil, true, "Two empty trees are identical"},
		{"one nil", NewNode(1), nil, false, "One empty, one non-empty"},
		{"single identical", NewNode(5), NewNode(5), true, "Single nodes with same value"},
		{"single different", NewNode(5), NewNode(3), false, "Single nodes with different values"},
		{"identical trees", buildTree1(), buildTree1(), true, "Identical structure and values"},
		{"different structure", buildTree1(), buildTree2(), false, "Same values, different structure"},
		{"different values", buildTree1(), buildTree3(), false, "Same structure, different values"},
		{"left vs right", buildLeftOnly(), buildRightOnly(), false, "One left-only, one right-only"},
	}

	for _, tc := range testCases {
		result := IsSameTree(tc.tree1, tc.tree2)
		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %t, Got: %t\n", tc.expected, result)
		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func buildTree1() *TreeNode {
	//   1
	//  / \
	// 2   3
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	return root
}

func buildTree2() *TreeNode {
	//   1
	//  /
	// 2
	//  \
	//   3
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Left.Right = NewNode(3)
	return root
}

func buildTree3() *TreeNode {
	//   1
	//  / \
	// 2   4
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(4)
	return root
}

func buildLeftOnly() *TreeNode {
	//   1
	//  /
	// 2
	root := NewNode(1)
	root.Left = NewNode(2)
	return root
}

func buildRightOnly() *TreeNode {
	// 1
	//  \
	//   2
	root := NewNode(1)
	root.Right = NewNode(2)
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

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
