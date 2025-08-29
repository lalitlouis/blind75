# Serialize and Deserialize Binary Tree - Algorithm Guide

## Problem Statement
Design an algorithm to serialize a binary tree to a string and deserialize it back to the original tree structure.

**Serialization:** Convert tree → string
**Deserialization:** Convert string → tree

## Examples

### Tree Structure:
```
    1
   / \
  2   3
     / \
    4   5
```

### Serialized String:
```
"1,2,null,null,3,4,null,null,5,null,null"
```

## Algorithm Approach
**Preorder Traversal with Null Markers**

### Why Preorder?
- Process current node first, then children
- Natural for reconstruction: create node, then build subtrees
- Maintains parent-child relationship during parsing

### Why Null Markers?
- Preserve tree structure information
- Distinguish between missing left/right children
- Enable accurate reconstruction

## Serialization Logic

### Process: Current → Left → Right
```go
func serialize(root *TreeNode) string {
    if root == nil {
        return "null"
    }
    
    return strconv.Itoa(root.Val) + "," + 
           serialize(root.Left) + "," + 
           serialize(root.Right)
}
```

### Step-by-step for example tree:
1. Visit 1: "1" + serialize(left) + serialize(right)
2. Visit 2: "2" + "null" + "null" = "2,null,null"
3. Visit 3: "3" + serialize(4) + serialize(5)
4. Visit 4: "4,null,null"
5. Visit 5: "5,null,null"
6. Result: "1,2,null,null,3,4,null,null,5,null,null"

## Deserialization Logic

### Parse and Build: Current → Left → Right
```go
func deserialize(data string) *TreeNode {
    values := strings.Split(data, ",")
    index := 0
    return buildTree(values, &index)
}

func buildTree(values []string, index *int) *TreeNode {
    if *index >= len(values) || values[*index] == "null" {
        *index++
        return nil
    }
    
    val, _ := strconv.Atoi(values[*index])
    *index++
    
    node := &TreeNode{Val: val}
    node.Left = buildTree(values, index)   // Build left subtree
    node.Right = buildTree(values, index)  // Build right subtree
    
    return node
}
```

### Key Insight: Index Tracking
- Use pointer to index to track position in array
- Each recursive call advances the index
- Maintains correct parsing order across all recursive calls

## Why This Order Works

### Information Flow: Top-down
- Parent creates node first
- Then recursively creates children
- Each level builds on the previous level's structure

### Parsing Dependencies:
- Current node value needed before building children
- Left subtree must be built before right subtree
- Index must advance sequentially through preorder sequence

## Implementation Pattern

```go
// Serialize: Process current, then recurse
func serialize(root *TreeNode) string {
    if root == nil {
        return "null"
    }
    return root.Val + "," + serialize(left) + "," + serialize(right)
}

// Deserialize: Create current, then recurse
func buildTree(values []string, index *int) *TreeNode {
    if values[*index] == "null" {
        *index++
        return nil
    }
    
    node := createNode(values[*index])  // Process current
    *index++
    node.Left = buildTree(values, index)   // Then left
    node.Right = buildTree(values, index)  // Then right
    return node
}
```

## Complexity
- **Time:** O(n) for both serialize and deserialize
- **Space:** O(n) for string storage and recursion stack

## Key Insights
1. **Preorder traversal** preserves reconstruction order
2. **Null markers** maintain structural information
3. **Index tracking** ensures correct parsing sequence
4. **Symmetric operations** - serialize and deserialize mirror each other

## Verification Strategy
**Roundtrip Testing:** serialize(deserialize(data)) should equal original data

This confirms both operations work correctly and preserve tree structure.