
## Same Tree

**Problem:** Determine if two binary trees are identical (same structure and values).

**Approach:** Simultaneous tree traversal
- Base cases: Both nil = true, one nil = false
- Compare current values and recurse on both subtrees

**Illustration:**
```
Tree 1:     Tree 2:     Result: true
   1           1
  / \         / \
 2   3       2   3

Tree 1:     Tree 2:     Result: false
   1           1
  /             \
 2               2
```

**Solution Pattern:**
```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    
    return p.Val == q.Val && 
           isSameTree(p.Left, q.Left) && 
           isSameTree(p.Right, q.Right)
}
```

**Time:** O(min(m,n)), **Space:** O(min(m,n))