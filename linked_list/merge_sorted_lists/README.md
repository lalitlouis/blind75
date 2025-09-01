# Merge Two Sorted Lists - Algorithm Guide

## Problem Statement
Given two sorted linked lists, merge them into one sorted list by splicing together the nodes of the two lists.

## Example
```
Input:  list1 = 1->2->4, list2 = 1->3->4
Output: 1->1->2->3->4->4
```

## Core Approach: Dummy Node Technique

### Why Use a Dummy Node?
- Eliminates edge case handling for empty result list
- Provides consistent starting point regardless of which list has smaller first element
- Simplifies pointer management

### Implementation
```go
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    dummy := &ListNode{}  // Create dummy starting node
    current := dummy      // Track current position
    
    // Compare and merge while both lists have nodes
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }
    
    // Attach remaining nodes (one list is exhausted)
    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }
    
    return dummy.Next  // Skip dummy node
}
```

## Step-by-step Execution
```
list1: 1->2->4, list2: 1->3->4

Step 1: Compare 1 vs 1, choose list1
Result: dummy->1, list1=2->4, list2=1->3->4

Step 2: Compare 2 vs 1, choose list2  
Result: dummy->1->1, list1=2->4, list2=3->4

Step 3: Compare 2 vs 3, choose list1
Result: dummy->1->1->2, list1=4, list2=3->4

Step 4: Compare 4 vs 3, choose list2
Result: dummy->1->1->2->3, list1=4, list2=4

Step 5: Compare 4 vs 4, choose list1
Result: dummy->1->1->2->3->4, list1=nil, list2=4

Step 6: Attach remaining list2
Final: dummy->1->1->2->3->4->4
```

## Common Mistakes

### 1. Wrong Comparison Direction
```go
// Wrong: Takes larger element first
if list1.Val >= list2.Val

// Correct: Takes smaller element first  
if list1.Val <= list2.Val
```

### 2. Wrong Return Value
```go
// Wrong: Returns current pointer (points to last node)
return current

// Correct: Returns first real node after dummy
return dummy.Next
```

### 3. Inefficient Remaining Node Handling
```go
// Inefficient: Loop through remaining nodes
while list1 != nil {
    current.Next = list1
    list1 = list1.Next
    current = current.Next
}

// Efficient: Attach entire remaining list
if list1 != nil {
    current.Next = list1
}
```

## Alternative: Recursive Approach
```go
func mergeTwoListsRecursive(list1, list2 *ListNode) *ListNode {
    if list1 == nil { return list2 }
    if list2 == nil { return list1 }
    
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoListsRecursive(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoListsRecursive(list1, list2.Next)
        return list2
    }
}
```

## Complexity
- **Time:** O(m + n) where m, n are lengths of the lists
- **Space:** O(1) iterative, O(m + n) recursive (call stack)

## Key Patterns
1. **Dummy node technique** - eliminates edge case complexity
2. **Two-pointer comparison** - advance the chosen pointer
3. **Remaining attachment** - handle leftover nodes efficiently

This merging pattern appears in many problems involving sorted data structures.