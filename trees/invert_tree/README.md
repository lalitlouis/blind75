## Invert Binary Tree

**Problem:** Mirror a binary tree by swapping left and right subtrees at every node.

**Approach:** Recursive subtree swapping
- Base case: Nil node returns nil
- Swap left and right children, then recurse

**Illustration:**
```
Original:      Inverted:
     4             4
   /   \         /   \
  2     7       7     2
 / \   / \     / \   / \
1   3 6   9   9   6 3   1
```

**Solution Pattern:**
```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    // Swap children first
    root.Left, root.Right = root.Right, root.Left
    
    // Then recurse on swapped children
    invertTree(root.Left)
    invertTree(root.Right)
    
    return root
}
```

**Time:** O(n), **Space:** O(h)