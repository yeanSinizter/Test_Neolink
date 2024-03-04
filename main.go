package main

import (
	"fmt"
)

var base = map[string]int{
	"A": 1,
	"B": 5,
	"Z": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"R": 1000,
}

func calculator(alienNumeral string) {
	splitAlienReverse := reverseString(alienNumeral)
	valuesOfChar := make([]int, len(splitAlienReverse))

	for i, char := range splitAlienReverse {
		valuesOfChar[i] = base[string(char)]
	}

	var summary int

	for i := 0; i < len(valuesOfChar); i++ {
		current := valuesOfChar[i]

		if i+1 < len(valuesOfChar) && current > valuesOfChar[i+1] {
			summary += (current - valuesOfChar[i+1])
			i++ // skip the next character since it's already considered
		} else {
			summary += current
		}
	}

	fmt.Println("summary =>", summary)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	calculator("RCRZCAB")
}
