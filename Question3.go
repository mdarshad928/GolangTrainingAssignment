package main

import (
	"fmt"
	"math"
	"strings"
)

/* 3. Implement the function words2num() which takes a string and returns the number. */

func words2num(word string) int32 {
	var answer int32
	var dictionary = map[string]int8{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	list := strings.Fields(word)
	multiplier := len(list) - 1
	for _, digitword := range list {
		answer += (int32(math.Pow10(multiplier))) * int32(dictionary[digitword])
		multiplier--
	}
	return answer
}

func main() {
	fmt.Println(words2num("one two three"))
}
