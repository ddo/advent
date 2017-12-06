package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func readInput(input string) []int {
	inputs := strings.Split(input, "\t")

	// str to int
	banks := []int{}
	var tmp int
	var err error
	for i := 0; i < len(inputs); i++ {
		tmp, err = strconv.Atoi(inputs[i])
		if err != nil {
			return nil
		}
		banks = append(banks, tmp)
	}
	return banks
}

// Day6 .
func Day6(input string) (step int) {
	banks := readInput(input)
	fmt.Println(len(banks), "banks:", banks)

	revisions, _ := process(banks)

	step = len(revisions)
	fmt.Println("step:", step)
	return
}

func process(banks []int) (revisions [][]int, firstRevisionIndex int) {
	length := len(banks)
	var maxI int
	var pool int
	revisions = [][]int{copyArr(banks)}
	for {
		maxI = max(banks)

		// calculate
		pool = banks[maxI]
		banks[maxI] = 0

		// move blocks
		for i := maxI + 1; i < maxI+pool+1; i++ {
			banks[i%length]++
		}

		firstRevisionIndex = exists(revisions, banks)
		if firstRevisionIndex >= 0 {
			return
		}

		revisions = append(revisions, copyArr(banks))
	}
}

func copyArr(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

func max(arr []int) (max int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return
}

func exists(revisions [][]int, rev []int) int {
	for i := 0; i < len(revisions); i++ {
		if sameArr(revisions[i], rev) {
			return i
		}
	}
	return -1
}

func sameArr(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Bonus .
func Bonus(input string) (cycle int) {
	banks := readInput(input)
	fmt.Println(len(banks), "banks:", banks)

	revisions, firstRevisionIndex := process(banks)

	cycle = len(revisions) - firstRevisionIndex
	fmt.Println("cycle:", cycle)
	return
}
