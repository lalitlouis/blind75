# Tree Traversals and Kth Smallest Element - Algorithm Guide

## Problem Statement
Implement the three main tree traversal methods and use inorder traversal to find the kth smallest element in a BST.

## Tree Traversal Types

### 1. Preorder Traversal
**Order:** Root → Left → Right
**Use Cases:** Tree copying, expression evaluation, serialization

### 2. Inorder Traversal  
**Order:** Left → Root → Right
**Use Cases:** BST sorted output, expression parsing

### 3. Postorder Traversal
**Order:** Left → Right → Root
**Use Cases:** Tree deletion, calculating directory sizes

## Algorithm Implementations

### Preorder Traversal
```go
func PreorderTraversal(node *TreeNode, values *[]int) {
    if node == nil {
        return
    }
    
    *values = append(*values, node.Val)    // Process root
    PreorderTraversal(node.Left, values)   // Then left
    PreorderTraversal(node.Right, values)  // Then right
}
```

### Inorder Traversal
```go
func InOrderTraversal(node *TreeNode, values *[]int) {
    if node == nil {
        return
    }
    
    InOrderTraversal(node.Left, values)    // First left
    *values = append(*values, node.Val)    // Then root
    InOrderTraversal(node.Right, values)   // Then right
}
```

### Postorder Traversal
```go
func PostOrderTraversal(node *TreeNode, values *[]int) {
    if node == nil {
        return
    }
    
    PostOrderTraversal(node.Left, values)   // First left
    PostOrderTraversal(node.Right, values)  // Then right
    *values = append(*values, node.Val)     // Finally root
}
```

## Visual Example

```
Tree:
    1
   / \
  2   3
 / \
4   5

Preorder:  [1, 2, 4, 5, 3]  (Root → Left → Right)
Inorder:   [4, 2, 5, 1, 3]  (Left → Root → Right)  
Postorder: [4, 5, 2, 3, 1]  (Left → Right → Root)
```

## Kth Smallest Element in BST

### Key Insight
Inorder traversal of a BST produces values in sorted (ascending) order.

### Implementation
```go
func KthSmallestElement(node *TreeNode, k int) int {
    values := []int{}
    InOrderTraversal(node, &values)
    return values[k-1]  // k is 1-indexed
}
```

### Example
```
BST:     8
       /   \
      3    10
     / \     \
    1   6    14
       / \   /
      4   7 13

Inorder: [1, 3, 4, 6, 7, 8, 10, 13, 14]
3rd smallest: values[2] = 4
```

## Processing Order Analysis

### Why Different Orders Matter

**Preorder** - Good for:
- Creating copies (need parent before children)
- Serialization (structure preservation)
- Prefix expression evaluation

**Inorder** - Good for:
- BST sorted output
- Expression parsing (infix notation)
- Range queries in BST

**Postorder** - Good for:
- Safe deletion (children before parent)
- Calculating subtree properties
- Postfix expression evaluation

## Optimization for Kth Smallest

**Early Stopping Version:**
```go
func kthSmallestOptimized(root *TreeNode, k int) int {
    count := 0
    result := 0
    inorderEarly(root, k, &count, &result)
    return result
}

func inorderEarly(node *TreeNode, k int, count *int, result *int) {
    if node == nil || *count >= k {
        return
    }
    
    inorderEarly(node.Left, k, count, result)
    
    *count++
    if *count == k {
        *result = node.Val
        return
    }
    
    inorderEarly(node.Right, k, count, result)
}
```

## Complexity Analysis

### All Traversals:
- **Time:** O(n) - visit each node once
- **Space:** O(h + n) - recursion stack + output array

### Kth Smallest:
- **Basic version:** O(n) time, O(n) space
- **Optimized version:** O(k + h) time, O(h) space

## Key Insights

1. **BST Property:** Inorder traversal gives sorted sequence
2. **Recursion Pattern:** All use same structure, different processing order
3. **Memory Management:** Use pointers to avoid copying arrays in recursion
4. **Early Termination:** Can optimize kth smallest by stopping at k elements

The traversal order determines when node processing occurs relative to recursive calls, making each suitable for different algorithmic needs.