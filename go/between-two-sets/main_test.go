package main

import "testing"

var testTable = []struct {
	a        []int32
	b        []int32
	expected int32
}{
	{[]int32{2, 6}, []int32{24, 36}, 2},
	{[]int32{2, 4}, []int32{16, 32, 96}, 3},
	{[]int32{1}, []int32{1}, 1},
	{[]int32{1}, []int32{100}, 9},
	{[]int32{3}, []int32{2}, 0},
	{[]int32{4}, []int32{2, 4}, 0},
	{[]int32{2, 3}, []int32{12, 24}, 2},
	{[]int32{7, 3}, []int32{42, 84, 126}, 2},
	{[]int32{2, 4, 8}, []int32{64}, 4},
	{[]int32{1, 2, 3, 4, 5}, []int32{60, 120}, 1},
	{[]int32{1}, []int32{97}, 2},
	{[]int32{100}, []int32{100}, 1},
}

func Test_getTotal(t *testing.T) {
	for _, item := range testTable {
		result := getTotalX(item.a, item.b)
		if result != item.expected {
			t.Errorf("wanted: %d, got: %d. Item: %v", item.expected, result, item)
		}
	}
}
