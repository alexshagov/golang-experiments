package main

import "fmt"

func main() {
  data := make(map[string]string)

  data["key1"] = "value1"
  data["key2"] = "value2"

  for _, v := range data {
    fmt.Println("Value is: ", v)
  }
}
