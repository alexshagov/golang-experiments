package main

import "fmt"
import "strings"

func main() {
  var answer []string

  for i := 1; i < 100; i++ {
    if i % 3 == 0 {
      answer = append(answer, "Fizz")
    }

    if i % 5 == 0 {
      answer = append(answer, "Buzz")
    }

    if len(answer) == 0 {
      fmt.Println(i)
    } else {
      fmt.Println(strings.Join(answer, ""))
      answer = answer[:0]
    }
  }
}
