package main

import (
	"fmt"
)

/*

largestPrimeFactor(2) should return a number.
largestPrimeFactor(2) should return 2.
largestPrimeFactor(3) should return 3.
largestPrimeFactor(5) should return 5.
largestPrimeFactor(7) should return 7.
largestPrimeFactor(8) should return 2.
largestPrimeFactor(13195) should return 29.
largestPrimeFactor(600851475143) should return 6857.

*/

func main() {
	fmt.Println(largestPrimeFactor(2))
	fmt.Println(largestPrimeFactor(3))
	fmt.Println(largestPrimeFactor(5))
	fmt.Println(largestPrimeFactor(7))
	fmt.Println(largestPrimeFactor(13195))
	fmt.Println(largestPrimeFactor(600851475143))
}


func largestPrimeFactor(n int) int {
	var a []int

	for n % 2 == 0 {
		a = append(a, 2)
		n /= 2
	}

	f := 3

	for f * f <= n {
		if n % f == 0 {
			a = append(a, f)
			n /= f
		} else {
			f += 2
		}
	}

	if n != 1 {
		a = append(a, n)
	}

	return a[len(a) - 1]
}