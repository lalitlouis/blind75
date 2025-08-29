# Sum of Left Leaves - Algorithm Guide

## Problem Statement
Given the root of a binary tree, return the sum of all left leaves.

A **left leaf** is a leaf node that is the left child of another node.

## Definition
A node qualifies as a left leaf if:
1. It has no children (is a leaf)
2. It is the left child of its parent

## Examples

### Example 1:
```
    3
   / \
  9  20
    /  \
   15   7

Output: 24 (only node 9 is a left leaf)
```

### Example 2:
```
       1
     /   \
    2     3
   / \   / \
  4   5 6   7

Output: 10 (left leaves: 4 + 6 = 10)
```

## Algorithm Approach
**Recursive traversal with parent context**

The key insight is that when processing a node, you need to know whether it arrived as a left child or right child from its parent.

## Information Flow
**Top-down (parent to child)**
- Parent passes direction information to each child
- Child uses this context to determine if it should count itself
- Must process first (determine parameters) then recurse

## Step-by-step Logic

1. **Base case:** If node is null, return 0
2. **Check if left leaf:** If node is a leaf AND is a left child, return its value  
3. **Recurse:** Sum results from left subtree (marked as left) and right subtree (marked as right)

## Implementation Pattern

```go
func sumOfLeftLeaves(root *TreeNode) int {
    return helper(root, false) // Root is not a left child
}

func helper(node *TreeNode, isLeftChild bool) int {
    if node == nil {
        return 0
    }
    
    // If it's a leaf AND it's a left child
    if node.Left == nil && node.Right == nil && isLeftChild {
        return node.Val
    }
    
    // Recurse: left child gets true, right child gets false
    return helper(node.Left, true) + helper(node.Right, false)
}
```

## Visual Execution

```
Tree:        Context passed:
    1        helper(1, false)
   / \      /              \
  2   3    helper(2, true)  helper(3, false)  
 /       /
4       helper(4, true) → leaf + left child → return 4
```

## Why This Order?

**Must process first because:**
- Information flows DOWN from parent to child
- Parent must determine what context to pass before recursing
- Children need parent's information to make decisions

Compare to problems that recurse first:
- Those need child results to compute parent values (bottom-up flow)
- This problem needs parent context to evaluate children (top-down flow)

## Complexity
- **Time:** O(n) - visit each node once
- **Space:** O(h) - recursion depth equals tree height

## Key Insight
The direction (left vs right child) is context that must be passed down from parent, not computed from children. This determines the processing order.