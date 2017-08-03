package main

import "fmt"
import "time"

func main() {
  hour := time.Now().Hour()
  switch hour {
  case 14:
    fmt.Println("Dinner time!")
  default:
    fmt.Println("Hey, you! Let's work!")
  }
}
