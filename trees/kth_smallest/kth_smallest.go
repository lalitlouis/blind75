package main

import "fmt"

func main() {
	TestPreorderTraversal()
	TestInorderTraversal()
	TestPostOrderTraversal()
	TestKthSmallestElement()
}

func TestPreorderTraversal() {
	fmt.Println("=== PREORDER TRAVERSAL TESTS ===")

	testCases := []struct {
		name     string
		tree     *TreeNode
		expected []int
		desc     string
	}{
		{"empty tree", nil, []int{}, "Empty tree produces empty array"},
		{"single node", NewNode(1), []int{1}, "Single node"},
		{"example tree", buildExampleTree(), []int{3, 9, 20, 15, 7}, "Root -> Left -> Right"},
		{"left skewed", buildLeftSkewed(), []int{1, 2, 3}, "Left-skewed tree"},
		{"right skewed", buildRightSkewed(), []int{1, 2, 3}, "Right-skewed tree"},
		{"complete tree", buildCompleteTree(), []int{1, 2, 4, 5, 3, 6, 7}, "Complete binary tree"},
	}

	for _, tc := range testCases {
		values := []int{}
		PreorderTraversal(tc.tree, &values)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %v, Got: %v\n", tc.expected, values)
		if slicesEqual(values, tc.expected) {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func TestInorderTraversal() {
	fmt.Println("=== INORDER TRAVERSAL TESTS ===")

	testCases := []struct {
		name     string
		tree     *TreeNode
		expected []int
		desc     string
	}{
		{"empty tree", nil, []int{}, "Empty tree produces empty array"},
		{"single node", NewNode(1), []int{1}, "Single node"},
		{"example tree", buildExampleTree(), []int{9, 3, 15, 20, 7}, "Left -> Root -> Right"},
		{"BST tree", buildBSTTree(), []int{1, 3, 4, 6, 7, 8, 10, 13, 14}, "BST gives sorted order"},
		{"left skewed", buildLeftSkewed(), []int{3, 2, 1}, "Left-skewed tree"},
		{"right skewed", buildRightSkewed(), []int{1, 2, 3}, "Right-skewed tree"},
		{"complete tree", buildCompleteTree(), []int{4, 2, 5, 1, 6, 3, 7}, "Complete binary tree"},
	}

	for _, tc := range testCases {
		values := []int{}
		InOrderTraversal(tc.tree, &values)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %v, Got: %v\n", tc.expected, values)
		if slicesEqual(values, tc.expected) {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func TestPostOrderTraversal() {
	fmt.Println("=== POSTORDER TRAVERSAL TESTS ===")

	testCases := []struct {
		name     string
		tree     *TreeNode
		expected []int
		desc     string
	}{
		{"empty tree", nil, []int{}, "Empty tree produces empty array"},
		{"single node", NewNode(1), []int{1}, "Single node"},
		{"example tree", buildExampleTree(), []int{9, 15, 7, 20, 3}, "Left -> Right -> Root"},
		{"left skewed", buildLeftSkewed(), []int{3, 2, 1}, "Left-skewed tree"},
		{"right skewed", buildRightSkewed(), []int{3, 2, 1}, "Right-skewed tree"},
		{"complete tree", buildCompleteTree(), []int{4, 5, 2, 6, 7, 3, 1}, "Complete binary tree"},
	}

	for _, tc := range testCases {
		values := []int{}
		PostOrderTraversal(tc.tree, &values)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %v, Got: %v\n", tc.expected, values)
		if slicesEqual(values, tc.expected) {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func TestKthSmallestElement() {
	fmt.Println("=== KTH SMALLEST ELEMENT TESTS ===")

	bst := buildBSTTree()

	testCases := []struct {
		name     string
		tree     *TreeNode
		k        int
		expected int
		desc     string
	}{
		{"1st smallest", bst, 1, 1, "Smallest element in BST"},
		{"8th element", bst, 8, 13, "8th smallest element"},    // 13, not 14
		{"5th smallest", bst, 5, 7, "Middle element"},          // 7, not 8
		{"last element", bst, 9, 14, "Largest element in BST"}, // 9th element is 14
		{"single node", NewNode(42), 1, 42, "Single node BST"},
	}

	for _, tc := range testCases {
		result := KthSmallestElement(tc.tree, tc.k)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %d, Got: %d\n", tc.expected, result)
		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

// Helper functions
func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
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

func buildBSTTree() *TreeNode {
	//       8
	//     /   \
	//    3    10
	//   / \     \
	//  1   6    14
	//     / \   /
	//    4   7 13
	root := NewNode(8)
	root.Left = NewNode(3)
	root.Right = NewNode(10)
	root.Left.Left = NewNode(1)
	root.Left.Right = NewNode(6)
	root.Left.Right.Left = NewNode(4)
	root.Left.Right.Right = NewNode(7)
	root.Right.Right = NewNode(14)
	root.Right.Right.Left = NewNode(13)
	return root
}

func buildLeftSkewed() *TreeNode {
	//   1
	//  /
	// 2
	///
	//3
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

// Your functions here...

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	*values = append(*values, node.Val)
	PreorderTraversal(node.Left, values)
	PreorderTraversal(node.Right, values)
}

func InOrderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	InOrderTraversal(node.Left, values)
	*values = append(*values, node.Val)
	InOrderTraversal(node.Right, values)
}

func PostOrderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	PostOrderTraversal(node.Left, values)
	PostOrderTraversal(node.Right, values)
	*values = append(*values, node.Val)
}

func KthSmallestElement(node *TreeNode, k int) int {
	values := []int{}
	InOrderTraversal(node, &values)
	return values[k-1]
}
