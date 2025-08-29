package main

import "fmt"

func main() {
	TestSumOfLeftLeaves()
}

func TestSumOfLeftLeaves() {
	testCases := []struct {
		name     string
		tree     *TreeNode
		expected int
		desc     string
	}{
		{"empty tree", nil, 0, "Nil tree has no left leaves"},
		{"single node", NewNode(1), 0, "Single node is not a left leaf"},
		{"example tree", buildExampleTree(), 24, "Only node 9 is a left leaf"},
		{"all left leaves", buildAllLeftLeaves(), 4, "Only node 4 is a left leaf"},
		{"no left leaves", buildNoLeftLeaves(), 0, "All leaves are right leaves"},
		{"mixed tree", buildMixedTree(), 13, "Left leaves: 4(4) + 9(9) = 13"},
		{"left skewed", buildLeftSkewed(), 3, "Only bottom node 3 is a left leaf"},
		{"right skewed", buildRightSkewed(), 0, "No left leaves in right-skewed tree"},
		{"complete tree", buildCompleteTree(), 10, "Left leaves: 4(4) + 6(6) = 10"},
	}

	for _, tc := range testCases {
		result := SumOfLeftLeaves(tc.tree, false) // Root is not a left child
		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)
		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func buildExampleTree() *TreeNode {
	//     3
	//    / \
	//   9  20
	//     /  \
	//    15   7
	// Left leaves: 9 (sum = 9)
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	return root
}

func buildAllLeftLeaves() *TreeNode {
	//   1
	//  / \
	// 2   5
	///     \
	//3       6
	///         \
	//4           7
	// Left leaves: 3, 4 (sum = 7)
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(5)
	root.Left.Left = NewNode(3)
	root.Right.Right = NewNode(6)
	root.Left.Left.Left = NewNode(4)
	root.Right.Right.Right = NewNode(7)
	return root
}

func buildNoLeftLeaves() *TreeNode {
	//   1
	//    \
	//     2
	//      \
	//       3
	// Left leaves: none (sum = 0)
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Right = NewNode(3)
	return root
}

func buildMixedTree() *TreeNode {
	//       1
	//     /   \
	//    2     3
	//   / \     \
	//  4   5     6
	//           /
	//          9
	// Left leaves: 4, 9 (sum = 13)
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Right = NewNode(6)
	root.Right.Right.Left = NewNode(9)
	return root
}

func buildLeftSkewed() *TreeNode {
	//   1
	//  /
	// 2
	///
	//3
	// Left leaves: 3 (sum = 3)
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Left.Left = NewNode(3)
	return root
}

func buildRightSkewed() *TreeNode {
	// 1
	//  \
	//   2
	//    \
	//     3
	// Left leaves: none (sum = 0)
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Right = NewNode(3)
	return root
}

func buildCompleteTree() *TreeNode {
	//       1
	//     /   \
	//    2     3
	//   / \   / \
	//  4   5 6   7
	// Left leaves: 4, 6 (sum = 10)
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)
	return root
}

// Your functions here...

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumOfLeftLeaves(node *TreeNode, isLeftChild bool) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil && isLeftChild {
		return node.Val
	}

	return SumOfLeftLeaves(node.Left, true) + SumOfLeftLeaves(node.Right, false)
}
