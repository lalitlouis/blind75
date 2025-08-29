package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	TestSerialize()
	TestDeSerialization()
}

func TestSerialize() {
	fmt.Println("=== SERIALIZE TESTS ===")

	testCases := []struct {
		name     string
		tree     *TreeNode
		expected string
		desc     string
	}{
		{"empty tree", nil, "null", "Null tree serializes to 'null'"},
		{"single node", NewNode(1), "1,null,null", "Single node with null children"},
		{"example tree", buildExampleTree(), "1,2,null,null,3,4,null,null,5,null,null", "Standard example tree"},
		{"left skewed", buildLeftSkewed(), "1,2,3,null,null,null,null", "Left-skewed tree"},
		{"right skewed", buildRightSkewed(), "1,null,2,null,3,null,null", "Right-skewed tree"},
		{"complete tree", buildCompleteTree(), "1,2,4,null,null,5,null,null,3,6,null,null,7,null,null", "Complete binary tree"},
	}

	for _, tc := range testCases {
		result := Serialize(tc.tree)
		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Expected: %s\n", tc.expected)
		fmt.Printf("Got: %s\n", result)
		if result == tc.expected {
			fmt.Println("PASS\n")
		} else {
			fmt.Println("FAIL\n")
		}
	}
}

func TestDeSerialization() {
	fmt.Println("=== DESERIALIZE TESTS ===")

	testCases := []struct {
		name string
		data string
		desc string
	}{
		{"empty tree", "null", "Deserialize null string"},
		{"single node", "1,null,null", "Deserialize single node"},
		{"example tree", "1,2,null,null,3,4,null,null,5,null,null", "Standard example tree"},
		{"left skewed", "1,2,3,null,null,null,null", "Left-skewed tree"},
		{"right skewed", "1,null,2,null,3,null,null", "Right-skewed tree"},
		{"complete tree", "1,2,4,null,null,5,null,null,3,6,null,null,7,null,null", "Complete binary tree"},
	}

	for _, tc := range testCases {
		tree := Deserialize(tc.data)
		serialized := Serialize(tree)

		fmt.Printf("Test: %s - %s\n", tc.name, tc.desc)
		fmt.Printf("Input: %s\n", tc.data)
		fmt.Printf("Roundtrip: %s\n", serialized)

		if serialized == tc.data {
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
	//     1
	//    / \
	//   2   3
	//      / \
	//     4   5
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Right.Left = NewNode(4)
	root.Right.Right = NewNode(5)
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

// Your functions here...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(node *TreeNode) string {
	if node == nil {
		return "null"
	}

	return strconv.Itoa(node.Val) + "," + Serialize(node.Left) + "," + Serialize(node.Right)

}

func Deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	index := 0
	return DeserializeHelper(values, &index)
}

func DeserializeHelper(values []string, index *int) *TreeNode {
	if *index >= len(values) || values[*index] == "null" {
		*index++
		return nil
	}

	val, _ := strconv.Atoi(values[*index])
	*index++

	node := &TreeNode{Val: val}
	node.Left = DeserializeHelper(values, index)
	node.Right = DeserializeHelper(values, index)

	return node
}
