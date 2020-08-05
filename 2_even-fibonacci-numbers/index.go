package main

import "fmt"

func main() {
	fmt.Println(fiboEvenSum(10))
	fmt.Println(fiboEvenSum(60))
	fmt.Println(fiboEvenSum(1000))
	fmt.Println(fiboEvenSum(100000))
	fmt.Println(fiboEvenSum(4000000))
}

func fiboEvenSum(n int) int {
	sum := 0
	Fa := 1
	Fb := 1

	for Fb <= n {
		Fc := Fa + Fb
		Fa = Fb
		Fb = Fc

		if (Fb % 2 == 0) {
			sum += Fb
		}
	}
	
	return sum
}