package day4

import (
	"fmt"
	"strings"
)

// Day4 .
func Day4(input string) (counter int) {
	// fmt.Println("input:", input)

	rows := strings.Split(input, "\n")
	for i := 0; i < len(rows); i++ {
		if !isDup(rows[i]) {
			counter++
		}
	}

	fmt.Println("counter:", counter)
	return
}

func isDup(str string) bool {
	items := strings.Split(str, " ")
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i] == items[j] {
				return true
			}
		}
	}

	return false
}

// Bonus .
func Bonus(input string) (counter int) {
	// fmt.Println("input:", input)

	rows := strings.Split(input, "\n")
	for i := 0; i < len(rows); i++ {
		if !isAnagram(rows[i]) {
			counter++
		}
	}

	fmt.Println("counter:", counter)
	return
}

func isAnagram(str string) bool {
	items := strings.Split(str, " ")

	// group words by length
	groupBylen := map[int][]string{}
	for i := 0; i < len(items); i++ {
		length := len(items[i])
		groupBylen[length] = append(groupBylen[length], items[i])
	}

	// check anagram
	for _, items := range groupBylen {

		strValues := map[int]string{}
		for i := 0; i < len(items); i++ {

			// calculate word value
			v := strToValue(items[i])

			// same value and same chars (not in order)
			_, ok := strValues[v]
			if ok && containSameChars(strValues[v], items[i]) {
				return true
			}

			strValues[v] = items[i]
		}
	}
	return false
}

func strToValue(str string) (v int) {
	for i := 0; i < len(str); i++ {
		v += int(str[i])
	}
	return
}

// not in order
func containSameChars(str1, str2 string) bool {
	for i := 0; i < len(str2); i++ {
		if !strings.Contains(str1, string(str2[i])) {
			return false
		}
	}
	return true
}
