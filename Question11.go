/* Q.11 Implement the function splitCopy() as per the details given below:

a) Function takes source file handle, destination file handle1, destination file handle2
	i) copies odd numbered lines to the file whose file handle is destination file handle1
	ii) copies even numbered lines to the file whose file handle is destination file handle2

b) Use the above function to copy odd and even numbered lines of file X into file Y and Z
	 respectively. Filenames X, Y and Z are given as command line arguments.*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func splitCopy(src, dest1, dest2 string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Open the destination files
	destFile1, err := os.Create(dest1)
	if err != nil {
		return err
	}
	defer destFile1.Close()

	destFile2, err := os.Create(dest2)
	if err != nil {
		return err
	}
	defer destFile2.Close()

	// Create a scanner for reading the source file
	scanner := bufio.NewScanner(srcFile)

	// Keep track of the line number
	lineNumber := 1

	// Loop through each line in the source file
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()

		// Write the line to the appropriate destination file
		if lineNumber%2 == 1 {
			_, err = destFile1.WriteString(line + "\n")
			if err != nil {
				return err
			}
		} else {
			_, err = destFile2.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}

		// Increment the line number
		lineNumber++
	}

	return nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <source file> <destination file 1> <destination file 2>")
		os.Exit(1)
	}

	src := os.Args[1]
	dest1 := os.Args[2]
	dest2 := os.Args[3]

	err := splitCopy(src, dest1, dest2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
