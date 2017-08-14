package main

import "fmt"
import "strconv"

func main() {
  var array [3]string
  
  for i := 0; i < 3; i++ {
    array[i] = strconv.Itoa(i);
  }

  fmt.Println(array)

  var mySlice = []string{"1", "2", "3"}
  var copiedSlice = make([]string, len(mySlice))
  copy(copiedSlice, mySlice)

  fmt.Println("MySlice len is:", len(mySlice))
  fmt.Println("MySlice cap is:", cap(mySlice))
  fmt.Println("Copied slice is:", copiedSlice)
  fmt.Println("Copied slice len is:", len(copiedSlice))
  fmt.Println("Copied slice cap is:", cap(copiedSlice))

  var twoD = make([][]int, 2)

  for i := 0; i < len(twoD); i++ {
    twoD[i] = make([]int, i + 1)
    for j := 0; j < i + 1; j++ {
      twoD[i][j] = i + j
    }
  }

  fmt.Println(twoD)

  // Creates second slice referencing to mySlice
  secondSlice := mySlice[:] 
  fmt.Println(secondSlice)
  mySlice[0] = "1+1"
  fmt.Println(secondSlice)

  // Creates other slice referencing to an array of original slice
  original := []int{1,2,3,4}
  other := original
  other[0] = 100
  fmt.Println(original)
  fmt.Println(other)

  // Appends one more integer to another slice, which causes creating of the new slice with different cap 
  other = append(original, 5)
  other[0] = 200
  fmt.Println(original)
  fmt.Println(other)
  fmt.Println("original slice cap: ", cap(original))
  fmt.Println("other slice cap: ", cap(other))
}
