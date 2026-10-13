package main

import "fmt"

func main() {
	arr := []int32{1, 2, 1, 2, 1, 3, 2}
	answer := sockMerchant(arr)
	fmt.Printf("answer: %d\n", answer)
}

func sockMerchant(arr []int32) int32 {
	var pairs int32
	seen := map[int32]bool{}
	for _, v := range arr {
		if seen[v] {
			pairs++
		}
		seen[v] = !seen[v]
	}
	return pairs
}
