package main

import (
	"fmt"
)

func main() {
	ar := []int32{1, 3, 2, 6, 1, 2}
	k := int32(3)
	l := int32(len(ar))
	answer := divisibleSumPairsImproved(l, k, ar)
	fmt.Printf("answer: %d\n", answer)
}

func divisibleSumPairs(n, k int32, ar []int32) int32 {
	var result int32
	for i, v_i := range ar {
		for j := i + 1; j < len(ar); j++ {
			if (v_i+ar[j])%k == 0 {
				result++
			}
		}
	}
	return result
}

func divisibleSumPairsImproved(n, target int32, ar []int32) int32 {
	// this will measure the frequency of remainders
	freq := map[int32]int32{}
	var count int32
	for _, v := range ar {
		remainder := v % target
		// want is the number that when added to the remainder equals the target
		// the extra mod is to handle the edge case where remainder is 0 and otherwise the want would
		// be target, which will never apear in the freq map as a remainder
		want := (target - remainder) % target
		// find out how often what we want has been seen
		count += freq[want]
		freq[remainder]++
	}
	return count
}
