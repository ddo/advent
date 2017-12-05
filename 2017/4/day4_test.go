package day4

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
	"aa bb cc dd ee":  1,
	"aa bb cc dd aa":  0,
	"aa bb cc dd aaa": 1,
}

func TestDay4(t *testing.T) {
	for k, v := range testCases {
		if Day4(k) != v {
			t.Error(testCases)
			return
		}
	}
}

func TestInput(t *testing.T) {
	Day4(input)
}

var testCasesBonus = map[string]int{
	"abcde fghij":              1,
	"abcde xyz ecdab":          0,
	"a ab abc abd abf abj":     1,
	"iiii oiii ooii oooi oooo": 1,
	"oiii ioii iioi iiio":      0,
	"www ww w wwwwww":          1,
	"www www w wwwww":          0,
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
