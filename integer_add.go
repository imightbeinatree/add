package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  a, _ := strconv.Atoi(os.Args[1])
  b, _ := strconv.Atoi(os.Args[2])
  result := strconv.Itoa(a + b)
  fmt.Println("{\"result\":\""+result+"\"}")
}


