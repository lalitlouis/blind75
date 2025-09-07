# Sliding Window Maximum

## Problem
Given array and window size `k`, return maximum value in each sliding window.

**Input:** `nums = [1,3,-1,-3,5,3,6,7], k = 3`  
**Output:** `[3,3,5,5,6,7]`

## Two Approaches

### Approach 1: Brute Force O(n×k)
```go
for i := 0; i <= len(nums)-k; i++ {
    result = append(result, max(nums[i:i+k]))
}
```
- **Pros:** Simple, easy to understand
- **Cons:** Inefficient for large k

### Approach 2: Deque O(n) ⭐
Use deque to maintain potential maximums efficiently.

## Deque Approach - Core Concept

**Deque stores indices in decreasing order of their values:**
- Front = current window maximum
- Back = potential future maximums
- Remove irrelevant elements as window slides

## Algorithm Steps

1. **Remove out-of-window indices** from front
2. **Remove smaller elements** from back (they'll never be maximum)
3. **Add current index** to back
4. **Record maximum** when window complete

## Code Template
```go
func SlidingWindowMax(nums []int, k int) []int {
    result := []int{}
    if len(nums) < k {
        return result
    }
    
    deque := make([]int, 0)  // stores indices
    
    for i := 0; i < len(nums); i++ {
        // Remove indices outside current window
        for len(deque) > 0 && deque[0] <= i-k {
            deque = deque[1:]  // remove from front
        }
        
        // Remove indices of smaller elements
        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
            deque = deque[:len(deque)-1]  // remove from back
        }
        
        // Add current index
        deque = append(deque, i)
        
        // Record result when window complete
        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }
    
    return result
}
```

## Key Insights

### Why Store Indices (Not Values)?
- **Window boundaries:** Need to check if element is still in window
- **Value comparison:** Use index to look up actual values
- **Position tracking:** Know where elements are in original array

### Why Decreasing Order?
- **Front is maximum:** Always contains index of largest element
- **Remove smaller:** Elements smaller than current will never be maximum
- **Maintain candidates:** Keep potential maximums for future windows

## Example Trace
```
nums = [1,3,-1,-3,5], k = 3

i=0: deque=[0], no result yet
i=1: 3>1, remove 0, deque=[1], no result yet  
i=2: deque=[1,2], result=[3] (first window [1,3,-1])
i=3: deque=[1,2,3], result=[3,3] 
i=4: 5>all, clear deque, deque=[4], result=[3,3,5]
```

## Time/Space Complexity
- **Time:** O(n) - each element added/removed at most once
- **Space:** O(k) - deque stores at most k elements

## When to Use Which Approach
- **O(n×k):** Small k, simple implementation needs
- **O(n) deque:** Large k, performance-critical, interviews

## Common Pitfalls
- Using `<` instead of `<=` for window boundary
- Storing values instead of indices in deque
- Starting loop from wrong position
- Forgetting to check if window is complete before recording result