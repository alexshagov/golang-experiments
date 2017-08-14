package main

import "fmt"

func makeAdder(x int) func(int) int {
  return func(y int) int {
    return x + y
  }
}

func main() {
  adder := makeAdder(5)
  fmt.Println(adder(2))
}
