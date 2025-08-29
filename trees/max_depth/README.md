## Maximum Depth of Binary Tree

**Problem:** Find the maximum depth (height) of a binary tree.

**Approach:** Recursive depth calculation
- Base case: Empty tree has depth 0
- Recursive case: 1 + max(left depth, right depth)

**Illustration:**
```
    3        depth = 3
   / \
  9  20      depth = 2
    /  \
   15   7    depth = 1
```

**Solution Pattern:**
```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    
    return 1 + max(leftDepth, rightDepth)
}
```

**Time:** O(n), **Space:** O(h) where h is tree height
