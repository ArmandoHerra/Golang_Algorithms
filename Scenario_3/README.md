# Scenario 3 - Fibonacci Sequence

## Problem Description

The Fibonacci sequence is a famous series of numbers where each number after the first two is the sum of the two preceding ones. The sequence often starts with 0 and 1 (though sometimes people start with 1 and 1):

```
F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2) for n >= 2
```

## Constraints

- Non-negative integer inputs (i.e., n >= 0).
- The function should handle both small (e.g., n = 0) and moderate-to-large n. However, extremely large n could cause performance or overflow concerns in some languages, so be mindful of that.
- We can consider iterative, recursive, or memoization/dynamic programming solutions.