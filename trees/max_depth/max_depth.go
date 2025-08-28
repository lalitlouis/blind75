package main

func main() {
	TestMaxDepth()
}

func TestMaxDepth() {

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
