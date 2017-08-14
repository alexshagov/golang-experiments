package main

import "fmt"

func main() {
  data := make(map[string]string)

  data["key"] = "value"

  fmt.Println(data["key"])

  _ , keyPresence := data["key2"] 

  fmt.Println(keyPresence)
}
