package main

import "fmt"

func main() {
	TestLCA()
}

func TestLCA() {
	testCases := []struct {
		name     string
		tree     *TreeNode
		p        int
		q        int
		expected int
		desc     string
	}{
		{"root LCA", buildBST(), 2, 8, 6, "LCA of nodes on different sides of root"},
		{"left subtree LCA", buildBST(), 0, 4, 2, "LCA within left subtree"},
		{"right subtree LCA", buildBST(), 7, 9, 8, "LCA within right subtree"},
		{"one node is ancestor", buildBST(), 2, 4, 2, "One node is ancestor of other"},
		{"same node", buildBST(), 4, 4, 4, "Both nodes are the same"},
		{"root is one node", buildBST(), 6, 8, 6, "Root is one of the nodes"},
		{"deep nodes", buildBST(), 3, 5, 4, "LCA of leaf nodes"},
		{"single node tree", NewNode(1), 1, 1, 1, "Tree with single node"},
		{"left edge case", buildBST(), 0, 2, 2, "Leftmost node with ancestor"},
		{"right edge case", buildBST(), 8, 9, 8, "Rightmost nodes"},
	}

	for _, tc := range testCases {
		result := LCA(tc.tree, tc.p, tc.q)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Searching for LCA of %d and %d\n", tc.p, tc.q)

		if result != nil {
			fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result.Val)
			if result.Val == tc.expected {
				fmt.Println("PASS\n")
			} else {
				fmt.Println("FAIL\n")
			}
		} else {
			fmt.Printf("Expected: %d, Got: nil\n", tc.expected)
			fmt.Println("FAIL\n")
		}
	}
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func buildBST() *TreeNode {
	//       6
	//     /   \
	//    2     8
	//   / \   / \
	//  0   4 7   9
	//     / \
	//    3   5
	root := NewNode(6)
	root.Left = NewNode(2)
	root.Right = NewNode(8)
	root.Left.Left = NewNode(0)
	root.Left.Right = NewNode(4)
	root.Left.Right.Left = NewNode(3)
	root.Left.Right.Right = NewNode(5)
	root.Right.Left = NewNode(7)
	root.Right.Right = NewNode(9)
	return root
}

// Your functions here...

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LCA(node *TreeNode, p, q int) *TreeNode {

	if node == nil {
		return nil
	}

	if p < node.Val && q < node.Val {
		return LCA(node.Left, p, q)
	} else if p > node.Val && q > node.Val {
		return LCA(node.Right, p, q)
	}

	return node
}
