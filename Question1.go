package main

import (
	"fmt"
	"math"
)

/* 1. Implement the function maxPower() that takes two parameters M and N and returns the largest integer K such that M <= N^K */

func maxPower(M int, N int) int {
	var i int
	var cond int
	for {
		i++
		cond = int(math.Pow(float64(N), float64(i)))
		if M <= cond {
			break
		}
	}

	return i
}

func main() {
	fmt.Println(maxPower(8000, 6))
}
