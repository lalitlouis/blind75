# Validate Binary Search Tree 
## Problem Statement
Given the root of a binary tree, determine if it is a valid Binary Search Tree (BST).

## Binary Search Tree Rules
1. All nodes in the left subtree must be **less than** the current node
2. All nodes in the right subtree must be **greater than** the current node
3. Both left and right subtrees must also be valid BSTs
4. No duplicate values allowed

## Algorithm Approach
**Range Validation with Recursion**

Each node must satisfy range constraints inherited from all its ancestors, not just its immediate parent.

## How It Works

### Step 1: Initialize with infinite range
```
validate(root, -∞, +∞)
```

### Step 2: For each node, check if value is within valid range
```
if node.Val <= minVal || node.Val >= maxVal:
    return false
```

### Step 3: Recursively validate children with updated ranges
```
Left child:  validate(left, minVal, node.Val)     // Upper bound narrows
Right child: validate(right, node.Val, maxVal)    // Lower bound narrows
```

## Visual Example

```
Tree:           Range Constraints:
    5           validate(5, -∞, +∞) ✓
   / \         /                    \
  3   8    validate(3,-∞,5) ✓    validate(8,5,+∞) ✓
 /   / \    /                    /              \
1   6   9  validate(1,-∞,3) ✓  validate(6,5,8) ✓  validate(9,8,+∞) ✓

Final result: true (all validations pass)
```

## Invalid Example

```
Tree:           Why It Fails:
    5           
   / \         
  3   7        
     / \      
    6   4      ← 4 is < 5, violates BST rule (should be > 5)

validate(4, 5, +∞) → 4 <= 5 → false
```

## Implementation Pattern

```go
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt, math.MaxInt)
}

func validate(node *TreeNode, minVal, maxVal int) bool {
    if node == nil {
        return true
    }
    
    if node.Val <= minVal || node.Val >= maxVal {
        return false
    }
    
    return validate(node.Left, minVal, node.Val) &&
           validate(node.Right, node.Val, maxVal)
}
```

## Key Insight
The range gets progressively **narrower** as you traverse deeper into the tree, ensuring all ancestor constraints are satisfied.

## Complexity
- **Time:** O(n) - visit each node once
- **Space:** O(h) - recursion depth equals tree height