package day1

import (
    "fmt"
)

// http://adventofcode.com/2017/day/1
func Day1(input string) (sum int) {
    fmt.Println("input:", input)
    package day3

import (
    "testing"
)

var testCases = map[int]int {
    // root
    1: 0,
    
    // vetical/horizontal
    2: 1,
    4: 1,
    6: 1,
    8: 1,
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
    3: 2,
    5: 2,
    7: 2,
    9: 2,
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
    
    // not special
    10: 3,
    12: 3,
    14: 3,
    16: 3,
    18: 3,
    20: 3,
    22: 3,
    24: 3,
    26: 5,
    27: 5,
    29: 5,
    30: 5,
    32: 5,
    33: 5,
    35: 5,
    36: 5,
    38: 5,
    39: 5,
    41: 5,
    42: 5,
    44: 5,
    45: 5,
    47: 5,
    48: 5,
    50: 7,
    51: 7,
    52: 7,
    54: 7,
    55: 7,
    56: 7,
    58: 7,
    59: 7,
    60: 7,
    62: 7,
    63: 7,
    64: 7,
    66: 7,
    67: 7,
    68: 7,
    70: 7,
    71: 7,
    72: 7,
    74: 7,
    75: 7,
    76: 7,
    78: 7,
    79: 7,
    80: 7,
    
    // cases
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
    // append last char for circular
    input += string(input[0])

    for i := 0; i < len(input) - 1; i++ {
        a, b := int(input[i])-48, int(input[i+1])-48
        if a == b {
            sum += a
        }
    }
    
    fmt.Println("sum:", sum)
    return
}

// http://adventofcode.com/2017/day/1 extra
func Day1Extra(input string) (sum int) {
    fmt.Println("input:", input)
    
    half := len(input)/2

    for i := 0; i < half; i++ {
        a, b := int(input[i])-48, int(input[i+half])-48
        if a == b {
            sum += a + a
        }
    }
    
    fmt.Println("sum:", sum)
    return
}