package day2

import (
	"io/ioutil"
	"testing"
)

var input string

func init() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input = string(data)
}

var testCases = map[string]int{
	`5	1	9	5
7	5	3
2	4	6	8`: 18,
}

func TestDay2(t *testing.T) {
	for k, v := range testCases {
		if Day2(k) != v {
			t.Error(testCases)
			return
		}
	}
}

func TestInput(t *testing.T) {
	Day2(input)
}

var testCasesBonus = map[string]int{
	`5	9	2	8
9	4	7	3
3	8	6	5`: 9,
}

func TestBonus(t *testing.T) {
	for k, v := range testCasesBonus {
		if Bonus(k) != v {
			t.Error(testCasesBonus)
			return
		}
	}
}

func TestInputBonus(t *testing.T) {
	Bonus(input)
}
