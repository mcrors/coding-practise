package main

import "testing"

var testTable = []struct {
	n        int32
	p        int32
	expected int32
}{
	{5, 3, 1},
	{6, 2, 1},              // sample 0
	{5, 4, 0},              // sample 1
	{1, 1, 0},              // one page book
	{2, 1, 0},              // first page
	{2, 2, 0},              // last page, even n
	{6, 6, 0},              // last page, even n
	{7, 7, 0},              // last page, odd n
	{6, 5, 1},              // back is shorter, even n
	{7, 4, 1},              // back is shorter, odd n
	{7, 3, 1},              // front is shorter, odd n
	{10, 5, 2},             // front is shorter, even n
	{10, 6, 2},             // both ways give the same count
	{100000, 1, 0},         // maximum n, first page
	{100000, 100000, 0},    // maximum n, last page
	{100000, 99999, 1},     // maximum n, one spread from the back
	{100000, 50000, 25000}, // maximum n, middle of the book
	{99999, 50000, 24999},  // maximum n, odd n, middle of the book
}

func Test_pageCount(t *testing.T) {
	for _, item := range testTable {
		actual := pageCount(item.n, item.p)
		if actual != item.expected {
			t.Errorf("wanted: %d. got: %d. n=%d. p=%d", item.expected, actual, item.n, item.p)
		}
	}
}
