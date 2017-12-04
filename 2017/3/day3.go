package day3

import (
	"fmt"
	"math"
)

type point struct {
	x     int
	y     int
	value int
}

// Day3 .
func Day3(input int) (step int) {
	fmt.Println("input:", input)

	if input <= 1 {
		step = 0
		return
	}

	sqrt := int(math.Ceil(math.Sqrt(float64(input))))

	width := sqrt

	// odd only
	if width%2 == 0 {
		width = width + 1
	}

	start := point{x: width / 2, y: width / 2}
	end := point{x: width - 1, y: width - 1}
	end.value = (end.x + 1) * (end.y + 1)

	// loop to find
	target := findSpiral(end, width, input)

	// manhattan distance
	step = abs(target.x-start.x) + abs(target.y-start.y)
	fmt.Println("step:", step)
	return
}

func abs(i int) int {
	if i < 0 {
		i = 0 - i
	}
	return i
}

func findSpiral(end point, width, input int) (target point) {
	target = end

	reverse := false
	for {
		reverse = !reverse

		// horizontal
		for i := 0; i < width-1; i++ {
			if target.value == input {
				return
			}

			target.value--
			if reverse {
				target.x--
			} else {
				target.x++
			}
		}

		// vertical
		for i := 0; i < width-1; i++ {
			if target.value == input {
				return
			}

			target.value--
			if reverse {
				target.y--
			} else {
				target.y++
			}
		}
	}
}
