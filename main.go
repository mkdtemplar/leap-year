package main

import (
	"fmt"
	"math"
)

func leapYearRule(yearDuration float32) []int {
	var currentYear = 2023
	var leapYear []int
	var difference float32

	for difference < 1 && currentYear <= 2043 {
		fraction := yearDuration - float32(math.Trunc(float64(yearDuration)))
		difference += fraction

		if difference >= 0.5 {
			difference -= 1
			leapYear = append(leapYear, currentYear)
		}
		currentYear++
	}

	return leapYear
}

func main() {
	fmt.Println(leapYearRule(365.26))
	fmt.Println(leapYearRule(400.45))
}
