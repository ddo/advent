package day3

import (
	"fmt"
	"math"
)

func Day3(input int) (step int) {
	fmt.Println("-------------------")
	fmt.Println("input:", input)

	if input <= 1 {
		step = 0
		return
	}

	// level
	sqrt := int(math.Sqrt(float64(input)))
	levelOdd := sqrt
	// odd only
	if sqrt%2 == 0 {
		levelOdd = sqrt - 1
	}
	level := levelOdd/2 + 1

	fmt.Println("sqrt:", sqrt)
	fmt.Println("levelOdd:", levelOdd)
	fmt.Println("level:", level)

	// diagonal
	if (input-(levelOdd*levelOdd))%(level*2) == 0 {
		fmt.Println("diagonal")
		if input == sqrt*sqrt {
			step = levelOdd - 1
		} else {
			step = levelOdd + 1
		}

		// vetical/horizontal
	} else if (input-(levelOdd*levelOdd+level))%(level*2) == 0 {
		fmt.Println("vetical/horizontal")
		step = level

		// not special
	} else {
		fmt.Println("not special")
		step = levelOdd
	}

	fmt.Println("step:", step)
	return
}
