# Two Sum - Algorithm

## Core Algorithm: HashMap One-Pass

**Key Insight:** Instead of checking all pairs (O(n²)), use a hashmap to find complements in O(1) time.

### Algorithm Steps:
1. Create empty hashmap: `number → index`
2. For each element at index `i`:
   - Calculate `complement = target - nums[i]`
   - If `complement` exists in hashmap → **Found pair!** Return indices
   - Else → Store `nums[i] → i` in hashmap
3. Continue until pair found

### Why It Works:
- Looking for: `a + b = target`
- When processing `a`, check if `b = target - a` was seen before
- HashMap provides O(1) lookup for the complement

### Walkthrough Example:
```
nums = [2, 7, 11, 15], target = 9

i=0: num=2, complement=7, hashmap={} → store 2→0, hashmap={2:0}
i=1: num=7, complement=2, hashmap={2:0} → found! return [0,1]
```

## Complexity:
- **Time:** O(n) - single pass
- **Space:** O(n) - hashmap storage