# Divisible Sum Pairs

## Problem

An array of integers and a positive integer `k` are given. Count the pairs `(i, j)` that meet two conditions.

1. `i < j`
2. `ar[i] + ar[j]` is divisible by `k`

## Example

`ar = [1, 2, 3, 4, 5, 6]` and `k = 5`.

Three pairs meet the criteria.

| Pair (indices) | Values | Sum |
| --- | --- | --- |
| `(0, 3)` | 1 and 4 | 5 |
| `(1, 2)` | 2 and 3 | 5 |
| `(3, 5)` | 4 and 6 | 10 |

## Function Description

Complete the function `divisibleSumPairs`.

`divisibleSumPairs` has these parameters:

- `int n`: the length of the array
- `int ar[n]`: an array of integers
- `int k`: the integer divisor

### Returns

- `int`: the number of pairs

## Input Format

The first line contains two space-separated integers, `n` and `k`.

The second line contains `n` space-separated integers. Each integer is a value of `ar[i]`.

## Constraints

- `2 <= n <= 100`
- `1 <= k <= 100`
- `1 <= ar[i] <= 100`

## Sample Input

```
STDIN           Function
-----           --------
6 3             n = 6, k = 3
1 3 2 6 1 2     ar = [1, 3, 2, 6, 1, 2]
```

## Sample Output

```
5
```

## Explanation

Here are the valid pairs when `k = 3`.

| Pair (indices) | Values | Sum |
| --- | --- | --- |
| `(0, 2)` | 1 and 2 | 3 |
| `(0, 5)` | 1 and 2 | 3 |
| `(1, 3)` | 3 and 6 | 9 |
| `(2, 4)` | 2 and 1 | 3 |
| `(4, 5)` | 1 and 2 | 3 |
