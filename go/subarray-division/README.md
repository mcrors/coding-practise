# Subarray Division

Two children, Lily and Ron, want to share a chocolate bar. Each square of the bar has an integer on it.

Lily decides to share a contiguous segment of the bar, selected such that:

- The length of the segment matches Ron's birth month, and
- The sum of the integers on the squares is equal to his birth day.

Determine how many ways she can divide the chocolate.

## Example

```
s = [2, 2, 1, 3, 2]
d = 4
m = 2
```

Lily wants to find segments summing to Ron's birth day, `d = 4`, with a length equal to his birth month, `m = 2`. Two segments meet her criteria: `[2, 2]` and `[1, 3]`.

## Function Description

Complete the `birthday` function.

`birthday` has the following parameters:

| Parameter | Description |
|-----------|-------------|
| `int s[n]` | The numbers on each of the squares of chocolate |
| `int d` | Ron's birth day |
| `int m` | Ron's birth month |

### Returns

- `int`: the number of ways the bar can be divided

## Input Format

1. The first line contains an integer `n`, the number of squares in the chocolate bar.
2. The second line contains `n` space-separated integers `s[i]`, the numbers on the chocolate squares, where `0 <= i < n`.
3. The third line contains two space-separated integers, `d` and `m`: Ron's birth day and his birth month.

## Constraints

- `1 <= n <= 100`
- `1 <= s[i] <= 5`, where `0 <= i < n`
- `1 <= d <= 31`
- `1 <= m <= 12`

## Sample Cases

### Sample Input 0

```
5
1 2 1 3 2
3 2
```

### Sample Output 0

```
2
```

### Explanation 0

Lily wants to give Ron `m = 2` squares summing to `d = 3`. Two segments meet the criteria: `[1, 2]` and `[2, 1]`.

### Sample Input 1

```
6
1 1 1 1 1 1
3 2
```

### Sample Output 1

```
0
```

### Explanation 1

Lily only wants to give Ron `m = 2` consecutive squares of chocolate whose integers sum to `d = 3`. Every pair of adjacent squares sums to `2`, so no segment satisfies these constraints. The answer is `0`.

### Sample Input 2

```
1
4
4 1
```

### Sample Output 2

```
1
```

### Explanation 2

Lily only wants to give Ron `m = 1` square of chocolate with an integer value of `d = 4`. The only square in the bar satisfies this constraint, so the answer is `1`.
