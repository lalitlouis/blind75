# Contains Duplicate - Algorithm

## Core Algorithm: HashMap Lookup

**Key Insight:** Use a hashmap to track elements we've seen. If we encounter an element that's already in the map, we've found a duplicate.

### Algorithm Steps:
1. Create empty hashmap to store seen elements
2. For each element in the array:
   - Check if element exists in hashmap
   - If **yes** → Duplicate found! Return `true`
   - If **no** → Add element to hashmap
3. If we finish the loop → No duplicates found, return `false`

### Why It Works:
- **Early termination:** Return immediately when first duplicate is found
- **O(1) lookups:** HashMap provides constant time checking
- **Space-time tradeoff:** Use O(n) space to achieve O(n) time

### Walkthrough Example:
```
nums = [1, 2, 3, 1]

i=0: num=1, hashmap={} → add 1, hashmap={1}
i=1: num=2, hashmap={1} → add 2, hashmap={1, 2}  
i=2: num=3, hashmap={1, 2} → add 3, hashmap={1, 2, 3}
i=3: num=1, hashmap={1, 2, 3} → found 1! Return true
```

```
nums = [1, 2, 3, 4]

i=0: num=1, hashmap={} → add 1, hashmap={1}
i=1: num=2, hashmap={1} → add 2, hashmap={1, 2}
i=2: num=3, hashmap={1, 2} → add 3, hashmap={1, 2, 3}
i=3: num=4, hashmap={1, 2, 3} → add 4, hashmap={1, 2, 3, 4}

End of loop → Return false
```

## Complexity:
- **Time:** O(n) - single pass through array
- **Space:** O(n) - hashmap can store up to n elements (worst case: no duplicates)