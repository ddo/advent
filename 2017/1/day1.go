package day1

import (
	"fmt"
)

// Day1 .
func Day1(input string) (sum int) {
	fmt.Println("input:", input)

	// append last char for circular
	input += string(input[0])

	for i := 0; i < len(input)-1; i++ {
		a, b := int(input[i])-48, int(input[i+1])-48
		if a == b {
			sum += a
		}
	}

	fmt.Println("sum:", sum)
	return
}

// Bonus .
func Bonus(input string) (sum int) {
	fmt.Println("input:", input)

	half := len(input) / 2

	for i := 0; i < half; i++ {
		a, b := int(input[i])-48, int(input[i+half])-48
		if a == b {
			sum += a + a
		}
	}

	fmt.Println("sum:", sum)
	return
}
