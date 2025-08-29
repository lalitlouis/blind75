package main

import "fmt"

func main() {
	TestLevelOrderTraversal()
}

func TestLevelOrderTraversal() {
	testCases := []struct {
		name     string
		tree     *TreeNode
		expected [][]int
		desc     string
	}{
		{"empty tree", nil, [][]int{}, "Nil tree returns empty result"},
		{"single node", NewNode(1), [][]int{{1}}, "Single node on level 0"},
		{"example tree", buildExampleTree(), [][]int{{3}, {9, 20}, {15, 7}}, "Standard example"},
		{"left skewed", buildLeftSkewed(), [][]int{{1}, {2}, {3}}, "Each level has one node"},
		{"right skewed", buildRightSkewed(), [][]int{{1}, {2}, {3}}, "Each level has one node"},
		{"complete tree", buildCompleteTree(), [][]int{{1}, {2, 3}, {4, 5, 6, 7}}, "Perfect binary tree"},
		{"unbalanced", buildUnbalanced(), [][]int{{1}, {2}, {3, 4}}, "Unbalanced tree"},
	}

	for _, tc := range testCases {
		result := LevelOrderTraversal(tc.tree)
		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Got: %v\n", result)

		if slicesEqual(result, tc.expected) {
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

func buildCompleteTree() *TreeNode {
	//       1
	//     /   \
	//    2     3
	//   / \   / \
	//  4   5 6   7
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)
	return root
}

func buildUnbalanced() *TreeNode {
	//   1
	//  /
	// 2
	///  \
	//3    4
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Left.Left = NewNode(3)
	root.Left.Right = NewNode(4)
	return root
}

func slicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Right: nil, Left: nil}
}

func LevelOrderTraversal(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel)
	}

	return result

}
