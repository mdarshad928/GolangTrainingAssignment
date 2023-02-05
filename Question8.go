/* 8. Write a program that takes 2 complex numbers and operation to be performed on them as command line arguments, and prints the result
as complex number.

Commandline arguments to the program are in the following format "complexnum1 binaryop complexnum2"

Note: Assume there is no embedded space in complex number argument.
ex:

3-4i + 7+2i
-15i - 3+4i
3-5i * 12
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[2], os.Args[1], os.Args[3])
	if len(os.Args) != 4 {
		fmt.Println("Pass Valid number of Arguments")
	} else if len(os.Args) == 4 {
		num1, err := strconv.ParseComplex(os.Args[1], 128)
		if err != nil {
			fmt.Println("Enter complex number as argument")
		}
		num2, err := strconv.ParseComplex(os.Args[3], 128)
		if err != nil {
			fmt.Println("Enter complex number as argument")
		}
		if os.Args[2] == "+" {
			fmt.Println(num1 + num2)
		} else if os.Args[2] == "-" {
			fmt.Println(num1 - num2)
		} else if os.Args[2] == "*" { // Passing * as operator in Bash shell is giving buggy result. Pass "*" in command line instead.
			fmt.Println(num1 * num2)
		} else {
			fmt.Println("Enter a valid Arithematic Operator")
		}
	}
}
