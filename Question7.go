/*7. Implement the function num2hex() which takes an unsighned integer and return its hexa equivalent as a string, Function should take an extra
argument which specifies whether the hexa digits should be in lower case (default) or upper case.

Using the function num2hex() implement main which takes an integer value and prints its hexa equivalent. If the number precedes with -u option,
then the hexa value should be printed in uppercase. Any other option other than -u should be considered invalid*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func num2hex(integer uint32, Case string) string {
	answer := ""
	var hexDigits string
	if Case == "u" {
		hexDigits = "0123456789ABCDEF"
	} else {
		hexDigits = "0123456789abcdef"
	}
	for integer > 0 {
		answer = string(hexDigits[integer%16]) + answer
		integer = integer / 16
	}
	return answer
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Enter command line arguments")
	} else if len(os.Args) == 2 {
		if integer, err := strconv.Atoi(os.Args[1]); err == nil {
			fmt.Println(num2hex(uint32(integer), "l"))
		}
	} else {
		if os.Args[1] == "-u" {
			if integer, err := strconv.Atoi(os.Args[2]); err == nil {
				fmt.Println(num2hex(uint32(integer), "u"))
			} else {
				fmt.Println("Invalid Argument")
			}
		}
	}
}
