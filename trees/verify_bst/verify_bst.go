package main

import "math"

func main() {
	TestVerifyBST()
}

func TestVerifyBST() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func ValidateBST(root *TreeNode) bool {
	return validateBSTHelper(root, math.MinInt, math.MaxInt)
}

func validateBSTHelper(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}

	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}

	return validateBSTHelper(node.Left, minVal, node.Val) && validateBSTHelper(node.Right, node.Val, maxVal)
}
