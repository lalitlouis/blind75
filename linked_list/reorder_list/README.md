# Reorder Linked List - Algorithm Guide

## Problem Statement
Reorder a linked list from L0→L1→...→Ln-1→Ln to L0→Ln→L1→Ln-1→L2→Ln-2→...

## Example
```
Input:  1→2→3→4→5
Output: 1→5→2→4→3
```

## Three-Step Algorithm

### Step 1: Find Middle and Split
```go
func FindMiddle(node *Node) *Node {
    slow := node
    fast := node.Next
    
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow  // Points to end of first half
}
```

**Split the list:**
```go
middleNode := FindMiddle(list)
secondHalf := middleNode.Next
middleNode.Next = nil  // Break connection
```

### Step 2: Reverse Second Half
```go
func Reverse(node *Node) *Node {
    var prev *Node
    curr := node
    
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
```

### Step 3: Merge Alternating
```go
func MergeLists(list1, list2 *Node) *Node {
    dummy := &Node{}
    list := dummy
    
    for list1 != nil && list2 != nil {
        list.Next = list1     // Take from first half
        list1 = list1.Next
        list = list.Next
        
        list.Next = list2     // Take from second half
        list2 = list2.Next
        list = list.Next
    }
    
    if list1 != nil {         // Attach remainder
        list.Next = list1
    }
    
    return dummy.Next
}
```

## Algorithm Walkthrough

**Input:** 1→2→3→4→5

**Step 1 - Find and Split:**
- Middle found at node 3
- First half: 1→2→3
- Second half: 4→5

**Step 2 - Reverse Second Half:**
- Original: 4→5
- Reversed: 5→4

**Step 3 - Merge Alternating:**
- Take 1 from first: 1
- Take 5 from second: 1→5
- Take 2 from first: 1→5→2
- Take 4 from second: 1→5→2→4
- Take 3 from first: 1→5→2→4→3

## Key Implementation Details

### Middle Finding Strategy
Using `fast = node.Next` initialization ensures proper splitting:
- **Odd length:** Middle stays with first half
- **Even length:** Split evenly

### Why Split is Essential
Without splitting, the reversed second half remains connected to the first half, creating cycles and incorrect merging.

### Alternating Pattern
Unlike value-based merging, this alternates by position regardless of node values. Always take one from each list in sequence.

## Edge Cases Handled
- **Single node:** Returns unchanged
- **Two nodes:** Natural alternation works
- **Odd length:** Extra node from first half attached at end
- **Even length:** Both halves consumed equally

## Complexity
- **Time:** O(n) - single pass for each operation
- **Space:** O(1) - only pointer manipulation

## Pattern Recognition
This problem combines three fundamental linked list techniques:
1. **Two-pointer traversal** (finding middle)
2. **Iterative reversal** (reversing second half)  
3. **Alternating merge** (reordering nodes)

The algorithm demonstrates how complex linked list operations can be decomposed into simpler, well-understood patterns.