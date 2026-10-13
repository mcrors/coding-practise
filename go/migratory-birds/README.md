# Migratory Birds

## Problem

An array holds bird sightings. Each element is a bird type id. Find the type id with the most sightings.

If more than one type has the maximum count, return the smallest id of those types.

## Example

```
arr = [1, 1, 2, 2, 3]
```

Types 1 and 2 each have two sightings. Type 3 has one sighting. The maximum count is 2. Type 1 is the lower of the two types with this count. The answer is 1.

## Function description

Complete the `migratoryBirds` function.

`migratoryBirds` has this parameter:

- `int arr[n]`: the types of the birds sighted.

It returns:

- `int`: the lowest type id of the most frequently sighted birds.

## Input format

The first line contains an integer `n`. This is the size of `arr`.

The second line contains `n` space-separated integers. Each integer is the type id of one sighted bird.

## Constraints

- 5 <= `n` <= 2 x 10^5
- Each type id is 1, 2, 3, 4, or 5.

## Sample input 0

```
6
1 4 4 4 5 3
```

## Sample output 0

```
4
```

### Explanation 0

| Type | Sightings |
| ---- | --------- |
| 1    | 1         |
| 2    | 0         |
| 3    | 1         |
| 4    | 3         |
| 5    | 1         |

Type 4 has the highest count. The answer is 4.

## Sample input 1

```
11
1 2 3 4 5 4 3 2 1 3 4
```

## Sample output 1

```
3
```

### Explanation 1

| Type | Sightings |
| ---- | --------- |
| 1    | 2         |
| 2    | 2         |
| 3    | 3         |
| 4    | 3         |
| 5    | 1         |

Types 3 and 4 each have a count of 3. Type 3 is the lower of the two. The answer is 3.

## Approach

1. Count the sightings of each type id.
2. Read the type ids in order, from 1 to 5.
3. Keep the first id that has the highest count.

The function reads the array one time. The time cost is O(n). The extra memory is O(1), because only five counters are used.
