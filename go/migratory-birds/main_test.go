package main

import "testing"

var testTable = []struct {
	arr      []int32
	expected int32
}{
	{[]int32{1, 4, 4, 4, 5, 3}, 4},
	// Sample input 1: types 3 and 4 tie, so the lower id wins
	{[]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}, 3},
	// Two types tie at the top
	{[]int32{1, 1, 2, 2, 3}, 1},
	// All sightings are one type
	{[]int32{5, 5, 5, 5, 5}, 5},
	// All five types tie
	{[]int32{1, 2, 3, 4, 5}, 1},
	// The tie is between the two highest ids
	{[]int32{4, 4, 5, 5, 1}, 4},
	// The lower id of the tie appears later in the input
	{[]int32{5, 5, 2, 2, 3}, 2},
	// The winner is type 1
	{[]int32{1, 1, 1, 2, 3, 4, 5}, 1},
	// The winner is type 5
	{[]int32{5, 5, 5, 1, 1, 2, 3, 4}, 5},
}

func Test_migratoryBirds(t *testing.T) {
	for _, item := range testTable {
		actual := migratoryBirds(item.arr)
		if actual != item.expected {
			t.Errorf("wanted: %d. got: %d", item.expected, actual)
		}
	}
}
