# Cycle Detection in Linked List - Algorithm Guide

## Problem Statement
Determine if a linked list contains a cycle (a node that points back to a previous node in the list).

## Example
```
Cycle exists:
3 -> 2 -> 0 -> -4
     ^         |
     |_________|

No cycle:
1 -> 2 -> 3 -> 4 -> null
```

## Algorithm: Floyd's Cycle Detection (Tortoise and Hare)

### Core Concept
Use two pointers moving at different speeds:
- **Slow pointer (tortoise):** moves 1 step at a time
- **Fast pointer (hare):** moves 2 steps at a time

If a cycle exists, the fast pointer will eventually catch up to the slow pointer.

### Implementation
```go
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    
    slow := head
    fast := head.Next
    
    for fast != nil && fast.Next != nil {
        if slow == fast {
            return true  // Cycle detected
        }
        slow = slow.Next      // Move 1 step
        fast = fast.Next.Next // Move 2 steps
    }
    
    return false  // No cycle found
}
```

### Why It Works
- In a linear list: fast pointer reaches the end first
- In a cyclic list: fast pointer eventually meets slow pointer
- The speed difference ensures they will meet within one cycle

### Edge Cases Handled
- **Empty list:** `head == nil`
- **Single node:** `head.Next == nil`
- **Two nodes:** Check `fast.Next != nil` before `fast.Next.Next`

## Common Mistakes

### 1. Incorrect Initialization
```go
// Can work but requires extra check in loop
slow := head
fast := head

// Cleaner - avoids initial equality
slow := head
fast := head.Next
```

### 2. Missing Null Checks
```go
// Wrong: Can cause null pointer panic
fast = fast.Next.Next

// Correct: Check both conditions
for fast != nil && fast.Next != nil {
    fast = fast.Next.Next
}
```

### 3. Wrong Loop Condition
```go
// Wrong: Misses the case where fast.Next is null
while fast != nil

// Correct: Must check both fast and fast.Next
while fast != nil && fast.Next != nil
```

## Complexity
- **Time:** O(n) - in worst case, traverse entire list once
- **Space:** O(1) - only use two pointers

## Key Insight
The relative speed difference guarantees that if a cycle exists, the fast pointer will "lap" the slow pointer exactly once before they meet. This happens because the fast pointer gains exactly one position on the slow pointer with each iteration.