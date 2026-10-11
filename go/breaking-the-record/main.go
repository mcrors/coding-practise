package main

import "fmt"

func main() {
	games := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	result := breakingRecords(games)
	fmt.Printf("Record highs beat: %d\n", result[0])
	fmt.Printf("Record lows beat: %d\n", result[1])
}

func breakingRecords(games []int32) [2]int32 {
	result := [2]int32{0, 0}
	high := games[0]
	low := games[0]
	for _, v := range games {
		if v > high {
			result[0]++
			high = v
		}
		if v < low {
			result[1]++
			low = v
		}
	}
	return result
}
