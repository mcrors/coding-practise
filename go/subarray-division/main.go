package main

import (
	"fmt"
)

func main() {
	s := []int32{1, 1, 1, 1}
	d := int32(2)
	m := int32(2)
	answer := birthday(s, d, m)
	fmt.Printf("answer: %d\n", answer)
}

// we track the window from the front backwards
// this way we remove the need for checks and the initial sum
// the window running out of space and the end of the slice become the
// same event
func lookback(s []int32, d, m int32) int32 {
	window := int(m)
	var total, result int32
	for i, v := range s {
		total += v
		if i >= window {
			total = total - s[i-window]
		}
		if i >= window-1 && total == d {
			result++
		}
	}
	return result
}

func birthday(s []int32, d, m int32) int32 {
	month := int(m)
	var result int32
	if month > len(s) {
		return result
	}

	total := sum(s[:month])
	for i := 0; i+month < len(s); i++ {
		if total == d {
			result++
		}

		total = total - s[i] + s[i+month]
	}

	if total == d {
		result++
	}

	return result
}

func sum(s []int32) int32 {
	total := int32(0)
	for _, v := range s {
		total += v
	}
	return total
}
