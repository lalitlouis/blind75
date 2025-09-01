package main

import "fmt"

func main() {
	TestInvertTree()
}

func TestInvertTree() {
	testCases := []struct {
		name     string
		tree     *TreeNode
		expected string
		desc     string
	}{
		{"empty tree", nil, "nil", "Empty tree remains empty"},
		{"single node", NewNode(1), "1", "Single node unchanged"},
		{"two nodes left", buildTwoLeft(), "1(nil,2)", "Left child becomes right"},
		{"two nodes right", buildTwoRight(), "1(2,nil)", "Right child becomes left"},
		{"example tree", buildExampleTree(), "4(7(9,6),2(3,1))", "Full example tree inverted"},
		{"symmetric tree", buildSymmetric(), "1(3(nil,7),3(7,nil))", "Symmetric tree inverted"},
	}

	for _, tc := range testCases {
		// Make a copy since InvertTree modifies the original
		treeCopy := copyTree(tc.tree)
		result := InvertTree(treeCopy)

		resultStr := treeToString(result)
		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %s, Got: %s\n", tc.expected, resultStr)

		if resultStr == tc.expected {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
	}
}

func buildTwoLeft() *TreeNode {
	// 1
	///
	//2
	root := NewNode(1)
	root.Left = NewNode(2)
	return root
}

func buildTwoRight() *TreeNode {
	// 1
	//  \
	//   2
	root := NewNode(1)
	root.Right = NewNode(2)
	return root
}

func buildExampleTree() *TreeNode {
	//     4
	//   /   \
	//  2     7
	// / \   / \
	//1   3 6   9
	root := NewNode(4)
	root.Left = NewNode(2)
	root.Right = NewNode(7)
	root.Left.Left = NewNode(1)
	root.Left.Right = NewNode(3)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(9)
	return root
}

func buildSymmetric() *TreeNode {
	//   1
	//  / \
	// 3   3
	//  \ /
	//   7 7
	root := NewNode(1)
	root.Left = NewNode(3)
	root.Right = NewNode(3)
	root.Left.Right = NewNode(7)
	root.Right.Left = NewNode(7)
	return root
}

// Helper functions for testing
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

func treeToString(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}
	return fmt.Sprintf("%d(%s,%s)", root.Val, treeToString(root.Left), treeToString(root.Right))
}

// Your functions here...

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Right: nil, Left: nil}
}

func InvertTree(p *TreeNode) *TreeNode {
	if p == nil {
		return nil
	}

	p.Left, p.Right = p.Right, p.Left

	InvertTree(p.Left)
	InvertTree(p.Right)

	return p
}
