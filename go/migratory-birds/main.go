package main

import "fmt"

func main() {
	arr := []int32{5, 5, 2, 2, 3}
	answer := migratoryBirds(arr)
	fmt.Printf("answer: %d\n", answer)
}

func migratoryBirds(arr []int32) int32 {
	var counts [6]int32
	winner := int32(1)
	for _, t := range arr {
		counts[t]++
		if (counts[t] > counts[winner]) || (counts[t] == counts[winner] && t < winner) {
			winner = t
		}
	}
	return winner
}
