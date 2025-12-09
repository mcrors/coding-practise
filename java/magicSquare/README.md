# Problem

We define a magic square to be an `n x n` matrix of distinct positive integers from `1`
to `n^2` where the sum of any row, column, or diagonal of length `n` is always equal to the
same number: the magic constant.

You will be given a `3 x 3` matrix  of integers in the inclusive range `[1, 9]`.
We can convert any digit `a` to any other digit `b` in the range `[1, 9]` at cost of `|a-b|`.
Given `s`, convert it into a magic square at minimal cost.

Print this cost on a new line.

*Note:* The resulting magic square must contain distinct integers in the inclusive range `[1, 9]`.

# Magic Constant
The magic constant can be gotten by summing all values in the square and dividing by the
number of rows in the square.

Number of rows = n

Sum of all numbers:
```
S = n^2(n^2 + 1) / 2
```

Magic constant:
```
M = S / n
```

For n = 3
```
M = (3^2(3^2 + 1) / 2) / 3
M = ((9*10) / 2) / 3)
M = (90 / 2) / 3
M = 45 / 3
M = 15
```

# Example

$s = [[5, 3, 4], [1, 5, 8], [6, 4, 2]]

The matrix looks like this:
```
5 3 4
1 5 8
6 4 2
```
We can convert it to the following magic square:
```
8 3 4
1 5 9
6 7 2
```
This took three replacements at a cost of `|5-8|+|8-9|+|4-7| = 7`.

# Function Description

Complete the formingMagicSquare function in the editor below.

formingMagicSquare has the following parameter(s):

int s[3][3]: a `3 x 3` array of integers

## Returns
int: the minimal total cost of converting the input square to a magic square

## Input Format
Each of the `3` lines contains three space-separated integers of row `s[i]`.

# Rough Work
There will be `2n + 2` lines or sums to do

so in a `3 x 3` square we will have `8` lines or arrays.
This is each row, each column, the diagonal and the anti-diagonal.

in the example we have
[5, 3, 4] -> col 1 = 12  3
[1, 5, 8] -> col 2 = 14  1
[6, 4 ,2] -> col 3 = 12  3
[5, 1, 6] -> row 1 = 12  3
[3, 5, 4] -> row 2 = 12  3
[4, 8, 2] -> row 3 = 14  1
[5, 5, 2] -> diagonal = 12  3
[4, 5, 6] -> anti-diagonal = 15

for each row
    for each column



recursion
tree
graph
hash
two pointers

5 numbers apear 3 times, 4 appear twice
The matrix looks like this:
```
5  3  4 -> 12
1  5  8 -> 14
6  4  2 -> 12
12 12 14
```
1, 2, 3, 4, 5, 6, 7, 8, 9
We can convert it to the following magic square:
```
8 3 4
1 5 9
6 7 2
```




