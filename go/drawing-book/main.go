package main

import "fmt"

func main() {
	n := int32(5)
	p := int32(3)
	answer := pageCount(n, p)
	fmt.Printf("answer: %d\n", answer)
}

func pageCount(n int32, p int32) int32 {
	frontWalk := p / 2
	backWalk := (n / 2) - (p / 2)
	if frontWalk < backWalk {
		return frontWalk
	}
	return backWalk
}
