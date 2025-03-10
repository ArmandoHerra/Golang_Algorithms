# Scenario 4 - Factorial

## Problem Description

A factorial of a positive integer n is defined as:

```
n!=n×(n−1)×(n−2)×⋯×2×1
```

By definition:
```
0!=1
```

and
```
1!=1
```

For example:
```
4!=4×3×2×1=24
5!=5×4×3×2×1=120
```
You’re tasked with creating a function factorial(n) that returns the factorial of n.

## Constraints

- n≥0; typically, we don’t define factorial for negative integers in standard mathematics.
- Factorials grow very fast and can exceed integer boundaries quickly. For languages like Go and Python, Python can handle arbitrarily large integers natively, but in Go, using built-in types (int or uint64) might lead to overflow if n is large. Keep that in mind for real-world usage.
- Time complexity for computing n!n! straightforwardly is O(n)O(n) if done iteratively, or can be done recursively with roughly the same complexity.