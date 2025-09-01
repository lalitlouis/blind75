# Search in Rotated Sorted Array - Algorithm

## Core Algorithm: Modified Binary Search

**Key Insight:** In a rotated sorted array, at least one half is always properly sorted. Identify the sorted half, check if target is in its range, then eliminate the appropriate half.

### Algorithm Steps:
1. Use standard binary search with `left`, `mid`, `right` pointers
2. At each step, determine which half is sorted:
   - If `nums[mid] >= nums[left]` → Left half is sorted
   - Else → Right half is sorted
3. Check if target is within the sorted half's range:
   - If yes → Search in that half
   - If no → Search in the other half
4. Repeat until target found or search space exhausted

### Why It Works:
- **Binary search**: O(log n) by eliminating half the search space each time
- **Rotation handling**: One half will always be properly sorted
- **Range checking**: Use sorted half's boundaries to determine target location

### Walkthrough Example:
```
nums = [4,5,6,7,0,1,2], target = 0

Step 1: left=0, mid=3, right=6, nums[mid]=7
- Left half [4,5,6,7] is sorted (7 >= 4)
- Target 0 not in range [4,7] → Go right

Step 2: left=4, mid=5, right=6, nums[mid]=1  
- Left half [0,1] is sorted (1 >= 0)
- Target 0 in range [0,1] → Go left

Step 3: left=4, mid=4, right=4, nums[mid]=0
- Found target at index 4!
```

## Complexity:
- **Time:** O(log n) - binary search
- **Space:** O(1) - only using pointers