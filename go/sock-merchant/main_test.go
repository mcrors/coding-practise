package main

import "testing"

var testTable = []struct {
	arr      []int32
	expected int32
}{
	{[]int32{1, 2, 1, 2, 1, 3, 2}, 2},
	// Sample input from the problem
	{[]int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
	// One sock has no pair
	{[]int32{1}, 0},
	// Exactly one pair
	{[]int32{7, 7}, 1},
	// Every sock is a different color
	{[]int32{1, 2, 3, 4, 5}, 0},
	// Every sock is the same color, even count
	{[]int32{4, 4, 4, 4, 4, 4}, 3},
	// Every sock is the same color, odd count leaves one over
	{[]int32{4, 4, 4, 4, 4, 4, 4}, 3},
	// Pairs are not next to each other in the input
	{[]int32{1, 2, 3, 1, 2, 3}, 3},
	// Colors at the limits of the constraint
	{[]int32{1, 100, 100, 1}, 2},
	// Three of one color gives one pair, not one and a half
	{[]int32{9, 9, 9, 8, 8}, 2},
}

func Test_sockMerchant(t *testing.T) {
	for _, item := range testTable {
		actual := sockMerchant(item.arr)
		if actual != item.expected {
			t.Errorf("wanted: %d. got: %d. arr: %v", item.expected, actual, item.arr)
		}
	}
}
