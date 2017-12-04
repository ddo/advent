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

func (p point) String() string {
	return fmt.Sprintf("(%v, %v) = %v", p.x, p.y, p.value)
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

const (
	right int = iota
	up
	left
	down
)

// Bonus .
func Bonus(input int) int {
	fmt.Println("input:", input)

	// generate grid
	grid := []point{}

	stepLen := 0
	direction := right
	target := point{x: 0, y: 0, value: 1}

	// 1st point
	grid = append(grid, target)

	for {
		// change step length
		stepLen++

		for i := 0; i < 2; i++ {
			// create points
			for j := 0; j < stepLen; j++ {

				switch direction {
				case right:
					target.x++
				case up:
					target.y--
				case left:
					target.x--
				case down:
					target.y++
				}

				target.value = calculatePointValue(grid, target)
				grid = append(grid, target)

				if target.value > input {
					fmt.Println("value:", target.value)
					return target.value
				}
			}

			// rotate
			direction = (direction + 1) % 4
		}
	}
}

func calculatePointValue(points []point, p point) (v int) {
	right := point{x: p.x + 1, y: p.y}
	rightUp := point{x: p.x + 1, y: p.y - 1}
	up := point{x: p.x, y: p.y - 1}
	leftUp := point{x: p.x - 1, y: p.y - 1}
	left := point{x: p.x - 1, y: p.y}
	leftDown := point{x: p.x - 1, y: p.y + 1}
	down := point{x: p.x, y: p.y + 1}
	rightDown := point{x: p.x + 1, y: p.y + 1}

	for i := len(points) - 1; i >= 0; i-- {
		if samePoint(points[i], right) ||
			samePoint(points[i], rightUp) ||
			samePoint(points[i], up) ||
			samePoint(points[i], leftUp) ||
			samePoint(points[i], left) ||
			samePoint(points[i], leftDown) ||
			samePoint(points[i], down) ||
			samePoint(points[i], rightDown) {
			v += points[i].value
		}
	}
	return
}

func samePoint(p1, p2 point) bool {
	return p1.x == p2.x && p1.y == p2.y
}
