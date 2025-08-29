# Lowest Common Ancestor of a Binary Search Tree - Algorithm Guide

## Problem Statement
Given a binary search tree (BST) and two nodes p and q, find the lowest common ancestor (LCA) of the two nodes.

The lowest common ancestor is the deepest node that has both p and q as descendants (a node can be a descendant of itself).

## Examples

### Example 1:
```
BST:        6
          /   \
         2     8
        / \   / \
       0   4 7   9
          / \
         3   5

Input: p = 2, q = 8
Output: 6 (root is LCA since 2 < 6 < 8)

Input: p = 2, q = 4  
Output: 2 (node 2 is ancestor of node 4)
```

## Algorithm Approach
**Leverage BST Property for Efficient Traversal**

The BST property allows us to determine the search direction without exploring both subtrees:
- All left descendants < root < all right descendants
- This gives us directional information for finding the LCA

## BST-Specific Logic

### Decision Tree:
1. **Both nodes < current node** → LCA must be in left subtree
2. **Both nodes > current node** → LCA must be in right subtree  
3. **Split case** (one left, one right, or equality) → current node is LCA

### Why This Works:
- If both nodes are on the same side, LCA cannot be current node
- If nodes are on different sides, current node is the split point (LCA)
- If one node equals current node, current node is the LCA

## Implementation

```go
func LCA(node *TreeNode, p, q int) *TreeNode {
    if node == nil {
        return nil
    }

    if p < node.Val && q < node.Val {
        return LCA(node.Left, p, q)   // Both in left subtree
    } else if p > node.Val && q > node.Val {
        return LCA(node.Right, p, q)  // Both in right subtree
    }

    return node  // Split case or equality - current node is LCA
}
```

## Step-by-step Execution

### Example: p = 2, q = 8 in BST above
1. **At node 6:** 2 < 6 and 8 > 6 → Split case → Return 6
2. **Result:** 6 is the LCA

### Example: p = 2, q = 4 in BST above  
1. **At node 6:** 2 < 6 and 4 < 6 → Both left → Go to node 2
2. **At node 2:** 2 = 2 and 4 > 2 → Split case → Return 2
3. **Result:** 2 is the LCA

## Edge Case Handling

### Equality Cases:
- If `p == node.Val` or `q == node.Val`, current node is LCA
- The split case condition handles this automatically

### Comparison Operators:
Using `<` and `>` (strict inequalities) prevents overlapping conditions and handles equality cases correctly.

## Complexity Analysis

### Time Complexity: O(log n) average, O(n) worst case
- **Average:** Balanced BST allows elimination of half the search space each step
- **Worst:** Skewed BST degenerates to linear search

### Space Complexity: O(log n) average, O(n) worst case
- Recursion stack depth equals tree height

## BST vs General Binary Tree

**BST Advantage:**
- Can determine direction without exploring both subtrees
- Time complexity: O(log n) average vs O(n) for general trees

**General Binary Tree:**
- Must explore both subtrees to find nodes
- Always O(n) time complexity

## Key Insights

1. **BST property enables direction-based decisions** without full tree traversal
2. **Split detection** - when nodes are on different sides of current node
3. **Equality handling** - when one node equals current node, current node is LCA
4. **Single path traversal** - unlike general trees, only one path needs to be explored

The BST structure transforms an otherwise complex tree problem into an efficient search problem.