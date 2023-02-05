package main

import (
	"fmt"
	"strings"
)

/*4. Implement the function capitalizeWords() which takes a string of words and returns a string containing the words such that first letter
of each word is capitalized.*/

func capitalizeWords(words string) string {
	return strings.Title(words)
}

func main() {
	fmt.Println(capitalizeWords("one two three"))
}
