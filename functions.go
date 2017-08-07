package main

import "fmt"

func arraySum(a []int) int {
  var sum int = 0

  for i := 0; i < len(a); i++ {
    sum += a[i]
  }

  return sum
}

func arraySumVariadic(numbers ...int) int {
  var sum int = 0

  for _, num := range numbers {
    sum += num
  }

  return sum
}

func main() {
  var array = []int{1, 2, 3}

  fmt.Println("#1 Sum of the array is: ", arraySum(array))
  fmt.Println("#2 Sum of the array is: ", arraySumVariadic(array...))
  fmt.Println("#3 Sum of the numbers 1,2,3 is: ", arraySumVariadic(1, 2, 3))
}
