package day1

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
	"1122":     3,
	"1111":     4,
	"1234":     0,
	"91212129": 9,
}

func TestDay1(t *testing.T) {
	for k, v := range testCases {
		if Day1(k) != v {
			t.Error(testCases)
			return
		}
	}
}

func TestInput(t *testing.T) {
	Day1(input)
}

var testCasesBonus = map[string]int{
	"1212":     6,
	"1221":     0,
	"123425":   4,
	"123123":   12,
	"12131415": 4,
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
