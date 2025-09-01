# Reverse Linked List - Algorithm Guide

## Problem Statement
Given the head of a singly linked list, reverse the list and return the new head.

## Example
```
Input:  1 -> 2 -> 3 -> 4 -> 5 -> null
Output: 5 -> 4 -> 3 -> 2 -> 1 -> null
```

## Linked List Structure
```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

## Approach 1: Iterative (Recommended)

### Core Concept
Use three pointers to track and reverse connections as you traverse.

### Implementation
```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode  // Previous node (starts as nil)
    curr := head        // Current node
    
    for curr != nil {
        next := curr.Next    // Store next before losing it
        curr.Next = prev     // Reverse the connection
        prev = curr          // Move prev forward
        curr = next          // Move curr forward
    }
    
    return prev  // prev is now the new head
}
```

### Step-by-step Execution
```
Initial: prev=nil, curr=1->2->3->null

Step 1: next=2, curr.Next=nil, prev=1, curr=2
Result: nil<-1  2->3->null

Step 2: next=3, curr.Next=1, prev=2, curr=3  
Result: nil<-1<-2  3->null

Step 3: next=nil, curr.Next=2, prev=3, curr=nil
Final: nil<-1<-2<-3
```

## Approach 2: Recursive

### Implementation
```go
func reverseListRecursive(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    newHead := reverseListRecursive(head.Next)
    head.Next.Next = head  // Reverse connection
    head.Next = nil        // Break forward link
    
    return newHead
}
```

### Key Insight
- Recursion handles the tail, then reverses connections on the way back up
- `head.Next.Next = head` creates the reverse link
- `head.Next = nil` prevents cycles

## Common Mistakes

### 1. Wrong Pointer Initialization
```go
// Wrong: Creates unnecessary node
prev := &ListNode{}

// Correct: Start with nil
var prev *ListNode
```

### 2. Incorrect Link Update
```go
// Wrong: Loses the connection
curr.Next = nil

// Correct: Points to previous
curr.Next = prev
```

### 3. Missing Next Storage
```go
// Wrong: Loses reference to rest of list
curr.Next = prev
curr = curr.Next  // This is now prev, not original next!

// Correct: Store next first
next := curr.Next
curr.Next = prev
curr = next
```

## Complexity
- **Time:** O(n) - visit each node once
- **Space:** O(1) iterative, O(n) recursive (call stack)

## Key Pattern
This problem introduces the fundamental linked list manipulation pattern:
1. Store next reference before modifying pointers
2. Update current node's pointer
3. Advance all tracking pointers

This pattern appears in many linked list problems.