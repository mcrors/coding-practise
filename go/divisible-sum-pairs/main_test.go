package main

import "testing"

var testTable = []struct {
	arr      []int32
	target   int32
	expected int32
}{
	{[]int32{1, 3, 2, 6, 1, 2}, 3, 5},
	// example from the problem statement
	{[]int32{1, 2, 3, 4, 5, 6}, 5, 3},
	// smallest array, one valid pair
	{[]int32{1, 2}, 3, 1},
	// smallest array, no valid pair
	{[]int32{1, 1}, 3, 0},
	// k = 1, so every pair is valid: 5 choose 2 = 10
	{[]int32{4, 7, 9, 2, 5}, 1, 10},
	// all values equal, every sum is 4: 4 choose 2 = 6
	{[]int32{2, 2, 2, 2}, 2, 6},
	// k is larger than every possible sum
	{[]int32{1, 2, 3}, 10, 0},
	// only sums that equal k exactly: (1,4) and (2,3)
	{[]int32{1, 2, 3, 4}, 5, 2},
	// every value is a multiple of k, so every pair is valid
	{[]int32{3, 6, 9, 12}, 3, 6},
	// remainders 1,1,2,3,3 with k = 4, valid pairs pair a 1 with a 3
	{[]int32{5, 9, 10, 7, 3}, 4, 4},
	// maximum values from the constraints
	{[]int32{100, 100, 100}, 100, 3},
	// order does not matter, a pair is counted once
	{[]int32{2, 1}, 3, 1},
}

func Test_divideSumPairs(t *testing.T) {
	for _, item := range testTable {
		n := int32(len(item.arr))
		actual := divisibleSumPairs(n, item.target, item.arr)
		if actual != item.expected {
			t.Errorf("wanted: %d. got: %d", item.expected, actual)
		}
	}
}
