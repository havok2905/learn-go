package main

import (
  "fmt" // A package in the Go standard library.
)

func main() {
  fmt.Println("hello world")
  beyondHello()
}

func beyondHello() {
  var x int
  x = 3
  y := 4 

  sum, prod := learnMultiple(x, y)
  fmt.Println("sum:", sum, "prod:", prod ) 
}

func learnMultiple(x, y int)(sum, prod int) {
  return x + y, x * y
}

func learnTypes() {
  str := "Learn Go" // string
  s2 := `A "raw" string literal
can include line breaks.` // raw string

  g := 'Σ' // non ascii
  f := 3.14195 // float
  c := 3 + 4i  // complex

  var u uint = 7 // float
  var pi float32 = 22. / 7 // float32

  n := byte('\n') // byte conversion

  var a4 [4]int // array
  a3 := [...]int{3, 1, 5} // inferred type array
  
  s3 := []int{4, 5, 9}  // slice
  s4 := make([]int, 4) // empty value slice
  var d2 [][]float64 // declared nested float slice
  bs := []byte("a slice") // infered slice

  s := []int{1, 2, 3} // inferred slice
  s = append(s, 4, 5, 6) // append values
  fmt.Println(s)
  
  s = append(s, []int{7, 8, 9}...) // append literal
  fmt.Println(s)

  _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d2, s3, a4, c, str, s2, g, f, u, pi, n, a3, s4, bs
}
