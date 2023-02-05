package main

import (
	"fmt"
	"strings"
)

/*5. Implement the function removeConsecutiveDuplicates() which takes two parameters, a string from which duplicates have to be removed and another
string that specifies characters to be removed. */

func removeConsecutiveDuplicates(input1, input2 string) string {
	var answer string
	input1 += "-"
	for i := 0; i < len(input1)-1; i++ {
		if !strings.ContainsRune(input2, rune(input1[i])) {
			answer += string(input1[i])
		} else {
			if i < len(input1) && !(rune(input1[i]) == rune(input1[i+1])) {
				answer += string(input1[i])
			} else if i < len(input1) && rune(input1[i]) == rune(input1[i+1]) {
				i++
			}
		}
	}

	return answer
}

func main() {
	fmt.Println(removeConsecutiveDuplicates("gooddeed", "od"))
}
