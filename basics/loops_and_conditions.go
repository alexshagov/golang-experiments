package main

import "fmt"

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }

  for i:= 0; i < 10000; i++ {
    fmt.Println(i)
    if i == 10 {
      break
    } else {
      fmt.Println("Ok, let's go further")
    }
  }
}
