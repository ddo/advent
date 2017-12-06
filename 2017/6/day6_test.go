package day6

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
	"0\t2\t7\t0": 5,
}

func TestDay6(t *testing.T) {
	for k, v := range testCases {
		if Day6(k) != v {
			t.Error(testCases)
			return
		}
	}
}

func TestInput(t *testing.T) {
	Day6(input)
}

var testCasesBonus = map[string]int{
	"0\t2\t7\t0": 4,
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
