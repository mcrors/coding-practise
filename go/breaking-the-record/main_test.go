package main

import "testing"

var testTable = []struct {
	input  []int32
	expect [2]int32
}{
	{[]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}, [2]int32{2, 4}},       // sample 0
	{[]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}, [2]int32{4, 0}}, // sample 1, only max breaks
	{[]int32{12, 24, 10, 24}, [2]int32{1, 1}},                      // example from the problem
	{[]int32{5}, [2]int32{0, 0}},                                   // single game sets the record, nothing to break
	{[]int32{7, 7, 7, 7}, [2]int32{0, 0}},                          // ties never break a record
	{[]int32{1, 2, 3, 4, 5}, [2]int32{4, 0}},                       // strictly increasing
	{[]int32{5, 4, 3, 2, 1}, [2]int32{0, 4}},                       // strictly decreasing
	{[]int32{0, 0, 0, 1}, [2]int32{1, 0}},                          // zero is a valid score
	{[]int32{100000000, 0, 100000000}, [2]int32{0, 1}},             // constraint max, tie at the top
}

func Test_breakingRecords(t *testing.T) {
	for _, item := range testTable {
		actual := breakingRecords(item.input)
		if item.expect != actual {
			t.Errorf("wanted: %v. got %v. item: %v", item.expect, actual, item)
		}
	}
}
