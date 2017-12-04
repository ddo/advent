package day2

import (
	"fmt"
	"strconv"
	"strings"
)

// Day2 .
func Day2(input string) (sum int) {
	fmt.Println("input:", input)
	rows := strings.Split(input, "\n")

	var items []string
	for i := 0; i < len(rows); i++ {
		items = strings.Split(rows[i], "\t")
		sum += findDiff(convertRow(items))
	}

	fmt.Println("sum:", sum)
	return
}

func convertRow(row []string) (items []int) {
	var item int
	for i := 0; i < len(row); i++ {
		item, _ = strconv.Atoi(row[i])
		items = append(items, item)
	}
	return
}

func findDiff(items []int) (diff int) {
	fmt.Println("items:", items)
	if len(items) == 0 {
		return
	}

	max := items[0]
	min := items[0]

	var item int
	for i := 0; i < len(items); i++ {
		item = items[i]

		// max
		if item > max {
			max = item
		}

		// min
		if item < min {
			min = item
		}
	}

	diff = max - min
	fmt.Println("diff:", diff)
	return
}

// Bonus .
func Bonus(input string) (sum int) {
	fmt.Println("input:", input)
	rows := strings.Split(input, "\n")

	var items []string
	for i := 0; i < len(rows); i++ {
		items = strings.Split(rows[i], "\t")
		sum += findEvenlyDiv(convertRow(items))
	}

	fmt.Println("sum:", sum)
	return
}

func findEvenlyDiv(items []int) (res int) {
	fmt.Println("items:", items)

	var a, b int
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			a, b = items[i], items[j]

			if dov(a, b) == 0 {
				return div(a, b)
			}
		}
	}
	return
}

func dov(a, b int) int {
	if a > b {
		return a % b
	}
	return b % a
}

func div(a, b int) int {
	if a > b {
		return a / b
	}
	return b / a
}
