package main

func main() {
	TestMaxPathSum()
}

func TestMaxPathSum() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func MaxPathSum(root *TreeNode) {
	globalMax := 0

	MaxPathSumHelper(root, globalMax)

}

func MaxPathSumHelper(root *TreeNode, globalSum int) int {
	if root == nil {
		return 0
	}

	leftMax := MaxPathSumHelper(root.Left, globalSum)
	rightMax := MaxPathSumHelper(root.Right, globalSum)
	leftMax = max(leftMax, 0)
	rightMax = max(rightMax, 0)

	globalSum = max(globalSum, leftMax+root.Val+rightMax)

	return root.Val + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
