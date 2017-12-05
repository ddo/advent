package day5

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
	`0
3
0
1
-3`: 5,
}

func TestDay5(t *testing.T) {
	for k, v := range testCases {
		if Day5(k) != v {
			t.Error(testCases)
			return
		}
	}
}

func TestInput(t *testing.T) {
	Day5(input)
}

var testCasesBonus = map[string]int{
	`0
3
0
1
-3`: 10,
}

func TestBonus(t *testing.T) {
	for k, v := range testCasesBonus {
		if Bonus(k) != v {
			t.Error(k, v)
			return
		}
	}
}

func TestInputBonus(t *testing.T) {
	Bonus(input)
}
