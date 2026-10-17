package main

import "fmt"

func main() {
	steps := int32(8)
	path := "DDUUUUDD"
	answer := countingValleys(steps, path)
	fmt.Printf("answer: %d\n", answer)
}

func countingValleys(steps int32, path string) int32 {
	var fromSeaLevel int32
	var valleys int32
	for _, s := range path {
		switch s {
		case 'U':
			fromSeaLevel++
		case 'D':
			fromSeaLevel--
		}
		if fromSeaLevel == 0 && s == 'U' {
			valleys++
		}
	}

	return valleys
}
