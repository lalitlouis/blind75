package main

import "fmt"

func main() {
	TestSubtreeOfAnotherTree()
}

func TestSubtreeOfAnotherTree() {
	testCases := []struct {
		name     string
		tree1    *TreeNode
		tree2    *TreeNode
		expected bool
		desc     string
	}{
		{"example true", buildMainTree(), buildSubTree(), true, "SubRoot exists in main tree"},
		{"not subtree", buildMainTree(), buildNonSubTree(), false, "SubRoot doesn't exist in main tree"},
		{"identical trees", buildSubTree(), buildSubTree(), true, "Identical trees - subtree of itself"},
		{"empty subtree", buildMainTree(), nil, true, "Empty tree is subtree of any tree"},
		{"empty main tree", nil, buildSubTree(), false, "No subtree can exist in empty tree"},
		{"single nodes match", NewNode(1), NewNode(1), true, "Single matching nodes"},
		{"single nodes different", NewNode(1), NewNode(2), false, "Single different nodes"},
		{"subtree at root", buildMainTree(), buildMainTree(), true, "Entire tree is subtree of itself"},
	}

	for _, tc := range testCases {
		result := IsSubTree(tc.tree1, tc.tree2)
		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %t, Got: %t\n", tc.expected, result)
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

func buildMainTree() *TreeNode {
	//     3
	//    / \
	//   4   5
	//  / \
	// 1   2
	root := NewNode(3)
	root.Left = NewNode(4)
	root.Right = NewNode(5)
	root.Left.Left = NewNode(1)
	root.Left.Right = NewNode(2)
	return root
}

func buildSubTree() *TreeNode {
	//   4
	//  / \
	// 1   2
	root := NewNode(4)
	root.Left = NewNode(1)
	root.Right = NewNode(2)
	return root
}

func buildNonSubTree() *TreeNode {
	//   4
	//  / \
	// 1   3
	root := NewNode(4)
	root.Left = NewNode(1)
	root.Right = NewNode(3) // Different from subtree
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSubTree(tree1, tree2 *TreeNode) bool {

	if tree1 == nil {
		return false
	}

	if tree2 == nil {
		return true
	}

	if isSameTree(tree1, tree2) {
		return true
	}

	return IsSubTree(tree1.Left, tree2) || IsSubTree(tree1.Right, tree2)

}

func isSameTree(tree1, tree2 *TreeNode) bool {

	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	return tree1.Val == tree2.Val && isSameTree(tree1.Left, tree2.Left) && isSameTree(tree1.Right, tree2.Right)
}
