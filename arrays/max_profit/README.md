# Best Time to Buy and Sell Stock - Algorithm

## Core Algorithm: Single Pass with Min Tracking

**Key Insight:** To maximize profit, always buy at the lowest price seen so far and check if selling at current price gives better profit.

### Algorithm Steps:
1. Initialize `minPrice = prices[0]` and `maxProfit = 0`
2. For each price in the array:
   - If `price < minPrice` → Update `minPrice = price`
   - Calculate `currentProfit = price - minPrice`
   - If `currentProfit > maxProfit` → Update `maxProfit = currentProfit`
3. Return `maxProfit`

### Why It Works:
- **Buy low, sell high:** Always track the minimum price encountered
- **Greedy approach:** At each step, calculate profit if we sell today
- **Constraint satisfied:** minPrice is always before current price (buy before sell)
- **Optimal substructure:** Best profit ending at day i = price[i] - min(price[0...i])

### Walkthrough Example:
```
prices = [7, 1, 5, 3, 6, 4]

i=0: price=7, min=7, profit=0, maxProfit=0
i=1: price=1, min=1, profit=0, maxProfit=0  
i=2: price=5, min=1, profit=4, maxProfit=4
i=3: price=3, min=1, profit=2, maxProfit=4
i=4: price=6, min=1, profit=5, maxProfit=5
i=5: price=4, min=1, profit=3, maxProfit=5

Result: 5 (buy at 1, sell at 6)
```

## Complexity:
- **Time:** O(n) - single pass through array
- **Space:** O(1) - only tracking min and maxProfit