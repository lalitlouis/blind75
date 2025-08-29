# Binary Tree Maximum Path Sum 
## Problem Statement
Given the root of a binary tree, return the maximum path sum of any non-empty path.

A **path** is any sequence of nodes from some starting node to any node in the tree along parent-child connections. The path must contain at least one node and does not need to go through the root.

## Examples

### Example 1:
```
Input: [1,2,3]
    1
   / \
  2   3

Output: 6 
Path: 2 → 1 → 3 (sum = 2 + 1 + 3 = 6)
```

### Example 2:
```
Input: [-10,9,20,null,null,15,7]
      -10
      /  \
     9   20
        /  \
       15   7

Output: 42
Path: 15 → 20 → 7 (sum = 15 + 20 + 7 = 42)
```

## Algorithm Approach
**Recursive path tracking with dual purpose function**

The key insight is that at each node, we need to track two different things:
1. **Maximum extendable path** - can continue upward to parent
2. **Maximum path in subtree** - includes paths that curve through current node

## Core Concepts

### Extendable vs Non-extendable Paths

```
At node with value 5:
     5
    / \
   3   4

Extendable paths (can go to parent):
- 3 → 5 (sum: 8)
- 5 → 4 (sum: 9)  ← Best extendable
- 5 alone (sum: 5)

Non-extendable path (curves, can't go to parent):
- 3 → 5 → 4 (sum: 12)  ← Best overall in subtree
```

### Algorithm Logic

```go
func maxPathHelper(node *TreeNode, globalMax *int) int {
    if node == nil {
        return 0
    }
    
    // Get best extendable paths from children (ignore negative)
    leftPath := max(0, maxPathHelper(node.Left, globalMax))
    rightPath := max(0, maxPathHelper(node.Right, globalMax))
    
    // Update global max with curved path through current node
    curvedPath := node.Val + leftPath + rightPath
    *globalMax = max(*globalMax, curvedPath)
    
    // Return best extendable path for parent
    return node.Val + max(leftPath, rightPath)
}
```

## Step-by-step Execution

### Tree: `[2, 1, 3]`
```
    2
   / \
  1   3
```

1. **At node 1:** return 1, globalMax = 1
2. **At node 3:** return 3, globalMax = 3  
3. **At node 2:**
   - leftPath = 1, rightPath = 3
   - curvedPath = 1 + 2 + 3 = 6
   - globalMax = max(3, 6) = 6
   - return 2 + max(1, 3) = 5

**Final answer:** 6

## Why Two Values?

The function returns the **extendable path** but tracks the **best overall path** globally because:
- Parents can only use straight paths from children
- The actual maximum might be a curved path that can't extend further

## Implementation Pattern

```go
func maxPathSum(root *TreeNode) int {
    globalMax := math.MinInt32
    maxPathHelper(root, &globalMax)
    return globalMax
}

func maxPathHelper(node *TreeNode, globalMax *int) int {
    if node == nil {
        return 0
    }
    
    left := max(0, maxPathHelper(node.Left, globalMax))
    right := max(0, maxPathHelper(node.Right, globalMax))
    
    *globalMax = max(*globalMax, node.Val + left + right)
    
    return node.Val + max(left, right)
}
```

## Complexity
- **Time:** O(n) - visit each node once
- **Space:** O(h) - recursion depth equals tree height

## Key Insight
At each node, the maximum path either:
1. Extends from a subtree through current node (extendable)
2. Curves through current node connecting both subtrees (non-extendable)

The algorithm tracks both possibilities to find the global maximum.