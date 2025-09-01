# Remove Nth Node From End - Algorithm Guide

## Problem Statement
Given a linked list, remove the nth node from the end and return the new head.

## Example
```
Input: [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Remove: node 4 (2nd from end)
```

## Core Approach: Two Pointers with Gap

### Algorithm Steps
1. Create dummy node pointing to head
2. Set both pointers to dummy
3. Move fast pointer n+1 steps ahead
4. Move both pointers until fast reaches end
5. Remove target node via slow.Next = slow.Next.Next

### Implementation
```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    slow := dummy
    fast := dummy
    
    // Move fast pointer n+1 steps ahead
    for i := 0; i <= n; i++ {
        if fast == nil {
            return head  // n too large
        }
        fast = fast.Next
    }
    
    // Move both until fast reaches end
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    
    // Remove the target node
    slow.Next = slow.Next.Next
    return dummy.Next
}
```

## Key Concepts Explained

### Why Dummy Node?
**Cannot use `dummy = head` because:**
- When removing head node, we need to return a new head
- Direct assignment creates circular references
- Dummy provides clean separation between old/new heads

**Example:**
```
Remove head from [1,2,3]:
dummy = head → points to node 1
After removal: still returns node 1 as head (wrong!)

dummy = &Node{Next: head} → separate placeholder
After removal: returns node 2 as new head (correct!)
```

### Pointer Safety Rules

**Check `fast == nil` when:**
- About to access `fast.Next` or `fast.Val`
- Moving one step: `fast = fast.Next`

**Check `fast != nil && fast.Next != nil` when:**
- About to access `fast.Next.Next` (two steps)
- Used in cycle detection

**In this problem:**
```go
for i := 0; i <= n; i++ {
    if fast == nil {        // Check BEFORE moving
        return head
    }
    fast = fast.Next        // Safe to move
}
```

### Why n+1 Steps?
We need `slow` to point to the node BEFORE the target:
```
[1,2,3,4,5], remove 2nd from end (node 4)

After n+1=3 steps: fast at node 4, slow at dummy
After moving together: fast at nil, slow at node 3
Remove: slow.Next (node 4) = slow.Next.Next (node 5)
```

## Common Mistakes

### 1. Wrong Dummy Usage
```go
// Wrong: Points to actual head
dummy = head

// Correct: Creates separate node
dummy = &ListNode{Next: head}
```

### 2. Assignment in Condition
```go
// Confusing: Assignment and check together
if fast = fast.Next; fast == nil

// Clear: Separate operations
fast = fast.Next
if fast == nil
```

### 3. Wrong Bounds Check Timing
```go
// Wrong: Check after moving (too late)
fast = fast.Next
if fast == nil

// Correct: Check before moving (prevents crash)
if fast == nil { return head }
fast = fast.Next
```

## Edge Cases Handled
- Empty list: `head == nil`
- Remove head: dummy node provides new head
- n larger than list: bounds checking prevents crash
- Single node: dummy technique works consistently

## Complexity
- **Time:** O(L) where L is list length
- **Space:** O(1) - only use constant extra pointers

The gap technique with dummy node provides a clean, safe solution for this pointer manipulation problem.