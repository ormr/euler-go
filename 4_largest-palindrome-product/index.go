package main

import "fmt"

func main() {
	fmt.Println(largestPalindromeProduct(2))
	fmt.Println(largestPalindromeProduct(3))
}

func reverseInteger(n int) int {
	reversedInt := 0

	for n > 0 {
		remainder := n % 10
		reversedInt *= 10
		reversedInt += remainder
		n /= 10
	}

	return reversedInt
}

func isPalindrome(n int) bool {
	return n == reverseInteger(n)
}

func largestPalindromeProduct(n int) int {
	largestPalindrome := 0
	maxDigit := 1

	for i := 0; i < n; i++ {
		maxDigit = maxDigit * 10
	}

	lowDigit := (maxDigit / 10)
	maxDigit = maxDigit - 1

	for i := maxDigit; i > lowDigit; i-- {
		for j := maxDigit; j > lowDigit; j-- {
			product := i * j

			if product > largestPalindrome {
				if (isPalindrome(product)) {
					largestPalindrome = product
				}
			} else if (product < largestPalindrome) {
				break
			}
		}
	}

	return largestPalindrome
}