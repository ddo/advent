package day3

import (
	"testing"
)

var testCases = map[int]int{
	// root
	1: 0,

	// vetical/horizontal
	2:  1,
	4:  1,
	6:  1,
	8:  1,
	11: 2,
	15: 2,
	19: 2,
	23: 2,
	28: 3,
	34: 3,
	40: 3,
	46: 3,
	53: 4,
	61: 4,
	69: 4,
	77: 4,

	// diagonal
	3:  2,
	5:  2,
	7:  2,
	9:  2,
	13: 4,
	17: 4,
	21: 4,
	25: 4,
	31: 6,
	37: 6,
	43: 6,
	49: 6,
	57: 8,
	65: 8,
	73: 8,
	81: 8,

	// random
	45:   4,
	1024: 31,
}

var input = 325489

func TestDay3(t *testing.T) {
	for k, v := range testCases {
		if Day3(k) != v {
			t.Error(k, v)
			return
		}
	}
}

func TestInput(t *testing.T) {
	Day3(input)
}

var testCasesBonus = map[int]int{
	1:   2,
	20:  23,
	50:  54,
	130: 133,
	134: 142,
	142: 147,
	500: 747,
	800: 806,
}

func TestBonus(t *testing.T) {
	for i := 0; i < len(testCasesBonus); i++ {
		for k, v := range testCasesBonus {
			if Bonus(k) != v {
				t.Error(k, v)
				return
			}
		}
	}
}

func TestInputBonus(t *testing.T) {
	Bonus(input)
}
