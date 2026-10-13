package main

import "fmt"

func main() {
	bill := []int32{3, 10, 2, 9}
	item := int32(1)
	charged := int32(12)
	bonAppetit(bill, item, charged)
}

// n = len(bill)
// k = item = the item that anna didn't eat
// b = charged = the amount  of money that brian charged anna
func bonAppetit(bill []int32, k int32, b int32) {
	var total int32
	for _, v := range bill {
		total += v
	}
	split := (total - bill[k]) / 2
	if split == b {
		fmt.Printf("Bon Appetit\n")
		return
	}
	fmt.Printf("%d\n", b-split)
}
