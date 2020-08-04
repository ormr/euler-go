package main

import "fmt"

func main() {
  fmt.Println(multiplesOf3and5(10))
  fmt.Println(multiplesOf3and5(49))
  fmt.Println(multiplesOf3and5(1000))
  fmt.Println(multiplesOf3and5(8456))
  fmt.Println(multiplesOf3and5(19564))
}

func multiplesOf3and5(number int) int {
  res := 0
  for i := 1; i < number; i++ {
    if (i % 3 == 0 || i % 5 == 0) {
      res += i
    }
  }

  return res
}