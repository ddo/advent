package day5

import (
	"fmt"
	"strconv"
	"strings"
)

// Day5 .
func Day5(input string) (step int) {
	inputs := strings.Split(input, "\n")

	// str to int
	instrctns := []int{}
	var tmp int
	var err error
	for i := 0; i < len(inputs); i++ {
		tmp, err = strconv.Atoi(inputs[i])
		if err != nil {
			return
		}
		instrctns = append(instrctns, tmp)
	}

	step = process(instrctns, func(offset int) int { return offset + 1 })
	fmt.Println("step:", step)
	return
}

func process(inputs []int, modifier func(int) int) (step int) {
	pos := 0
	var nextPos int
	for {
		// stop
		if pos >= len(inputs) {
			return step
		}

		// get next pos
		nextPos = inputs[pos] + pos

		// update current instruction
		inputs[pos] = modifier(inputs[pos])

		// update step
		step++

		// update pos
		pos = nextPos
	}
}

// Bonus .
func Bonus(input string) (step int) {
	inputs := strings.Split(input, "\n")

	// str to int
	instrctns := []int{}
	var tmp int
	var err error
	for i := 0; i < len(inputs); i++ {
		tmp, err = strconv.Atoi(inputs[i])
		if err != nil {
			return
		}
		instrctns = append(instrctns, tmp)
	}

	step = process(instrctns, func(offset int) int {
		if offset >= 3 {
			return offset - 1
		}
		return offset + 1
	})
	fmt.Println("step:", step)
	return
}
