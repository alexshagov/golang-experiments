package main

import "fmt"
import "strconv"

func main() {
  var array [3]string
  
  for i := 0; i < 3; i++ {
    array[i] = strconv.Itoa(i);
  }

  fmt.Println(array)

  var mySlice = []string{"1", "2"}
  var copiedSlice = make([]string, len(mySlice))
  copy(copiedSlice, mySlice)

  fmt.Println("MySlice len is:", len(mySlice))
  fmt.Println("Copied slice is:", copiedSlice)

  var twoD = make([][]int, 2)

  for i := 0; i < len(twoD); i++ {
    twoD[i] = make([]int, i + 1)
    for j := 0; j < i + 1; j++ {
      twoD[i][j] = i + j
    }
  }

  fmt.Println(twoD)
}
