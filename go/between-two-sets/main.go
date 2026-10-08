package main

import (
	"fmt"
)

func getTotalX(a []int32, b []int32) int32 {
	highest := gcdOfSlice(b)
	step := lcmOfSlice(a)
	var r int32
	for c := step; c <= highest; c += step {
		if highest%c == 0 {
			r++
		}
	}

	return r
}

func gcd(x int32, y int32) int32 {
	remainder := x % y
	for remainder != 0 {
		x = y
		y = remainder
		remainder = x % y
	}
	return y
}

func gcdOfSlice(nums []int32) int32 {
	result := nums[0]
	for i := range nums {
		if i+1 >= len(nums) {
			break
		}
		result = gcd(result, nums[i+1])
	}
	return result
}

func lcm(x int32, y int32) int32 {
	g := gcd(x, y)

	return (x / g) * y
}

func lcmOfSlice(nums []int32) int32 {
	result := nums[0]
	for i := range nums {
		if i+1 >= len(nums) {
			break
		}
		result = lcm(result, nums[i+1])
	}
	return result
}

func main() {
	a := []int32{1}
	b := []int32{100}

	r := getTotalX(a, b)
	fmt.Println(r)
}
