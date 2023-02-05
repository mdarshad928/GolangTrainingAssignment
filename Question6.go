package main

import "fmt"

/*6. Write a program that creates slices each of increasing size, starting with slice of 1 element. Store the first 120 prime numbers into these
slices as per the pattern given below.
		Slice-1 : first prime number
		Slice-2 : next 2 prime numbers
		Slice-3 : next 3 prime numbers
		Slice-4 : next 4 prime numbers
Having created, print contents of each slice one line per row. */
func nextPrime(num uint32) uint32 {
	if num == 1 || num == 2 {
		num++
		return num
	}
	num++
	for ; ; num++ {
		var j uint32
		flag := true
		for j = 2; j*j <= num; {
			if num%j == 0 {
				flag = false
				break
			}
			j++
		}
		if flag == true {
			return num
		}
	}
}
func primeNumSlice() {
	var num uint32
	num = 1
	for i := 1; i <= 15; i++ {
		result := make([]uint32, i)
		for j := 0; j < i; j++ {
			t := nextPrime(num)
			num = t
			result[j] = t
		}
		fmt.Println(result)
	}
}

func main() {
	primeNumSlice()
}
