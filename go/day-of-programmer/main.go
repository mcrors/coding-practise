package main

import "fmt"

func main() {

}

func dayOfProgrammer(year int32) string {
	if year == 1918 {
		return "26.09.1918"
	}
	var isLeap bool
	if year <= 1917 {
		isLeap = year%4 == 0
	} else {
		isLeap = year%400 == 0 || (year%4 == 0 && !(year%100 == 0))
	}
	day := 13
	if isLeap {
		day = 12
	}
	return fmt.Sprintf("%d.09.%d", day, year)
}
