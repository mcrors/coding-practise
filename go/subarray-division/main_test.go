package main

import "testing"

var testTable = []struct {
	s        []int32
	d        int32
	m        int32
	expected int32
}{
	{[]int32{2, 2, 1, 3, 2}, int32(4), int32(2), int32(2)},
	{[]int32{1, 2, 1, 3, 2}, int32(3), int32(2), int32(2)},    // sample 0
	{[]int32{1, 1, 1, 1, 1, 1}, int32(3), int32(2), int32(0)}, // sample 1: no matches
	{[]int32{4}, int32(4), int32(1), int32(1)},                // sample 2: single square
	{[]int32{1, 2}, int32(3), int32(3), int32(0)},             // m longer than the bar
	{[]int32{1, 2, 3}, int32(6), int32(3), int32(1)},          // segment is the whole bar
	{[]int32{2, 5, 2, 2}, int32(2), int32(1), int32(3)},       // m = 1, several matches
	{[]int32{1, 1, 1, 1}, int32(2), int32(2), int32(3)},
	{[]int32{1, 4, 3, 2, 2, 1, 5}, int32(4), int32(2), int32(1)}, // sliding window: only [2, 2] matches
}

func Test_birthday(t *testing.T) {
	for _, item := range testTable {
		actual := birthday(item.s, item.d, item.m)
		if actual != item.expected {
			t.Errorf("wanted: %d. got: %d", item.expected, actual)
		}
	}
}

func Test_lockback(t *testing.T) {
	for _, item := range testTable {
		actual := lookback(item.s, item.d, item.m)
		if actual != item.expected {
			t.Errorf("wanted: %d. got: %d", item.expected, actual)
		}
	}
}
