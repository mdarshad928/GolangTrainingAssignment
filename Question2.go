package main

import "fmt"

/* 2. Implement the function num2words() which takes number and returns the number in words as a string. */
func num2words(number int32) string {
	var digit int8
	var answer string
	var dictionary = map[int8]string{0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}
	for number > 0 {
		digit = int8(number % 10)
		answer = dictionary[digit] + " " + answer
		number /= 10
	}
	return answer[:len(answer)-1]
}

func main() {
	fmt.Println(num2words(123))
}
