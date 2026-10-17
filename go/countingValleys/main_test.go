package main

import "testing"

var testTable = []struct {
	steps    int32
	path     string
	expected int32
}{
	{int32(8), "UDDDUDUU", 1},
	{int32(8), "DDUUUUDD", 1},
	{int32(2), "UD", 0},
	{int32(2), "DU", 1},
	{int32(4), "DUDU", 2},
	{int32(4), "UDUD", 0},
	{int32(6), "DDDUUU", 1},
	{int32(6), "DUUDDU", 2},
	{int32(8), "UUDDDDUU", 1},
	{int32(10), "UDDUUDDUDU", 3},
	{int32(12), "DDUUDDUDUUUD", 2},
}

func Test_countingValleys(t *testing.T) {
	for _, item := range testTable {
		t.Run(item.path, func(t *testing.T) {
			actual := countingValleys(item.steps, item.path)
			if actual != item.expected {
				t.Errorf("countingValleys(%d, %q) = %d, want %d", item.steps, item.path, actual, item.expected)
			}
		})
	}
}
