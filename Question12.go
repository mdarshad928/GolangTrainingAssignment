/*12. Write a program that reads words from console, and at the end of the input, prints words of
      same length. Output only those words which have at least 4 characters.

			Output should be in the following format:

			4 : 4-letter words in alphabetical order
			5 : 5-letter words in alphabetical order
			6 : 6-letter words in alphabetical order
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	//	input := io.Reader(os.Stdin)
	//	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	lengthDict := map[int][]string{}
	fmt.Println("Enter the words and Press Ctrl+D when Done")
	for scanner.Scan() {
		word := scanner.Text()
		if value, ok := lengthDict[len(word)]; ok {
			lengthDict[len(word)] = append(value, word)
		} else {
			lengthDict[len(word)] = []string{word}
		}
	}
	for key, val := range lengthDict {
		fmt.Println("Length:", key, "words are", val)
	}
}
