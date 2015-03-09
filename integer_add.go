package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  a := strconv.Atoi(os.Args[1])
  b := strconv.Atoi(os.Args[2])
  result := strconv.Itoa(a + b)
  fmt.Println("{\"result\":\""+result+"\"}")
}


