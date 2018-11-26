package main

// solution for: https://tour.golang.org/moretypes/23

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  a := make(map[string]int)
  for _, word := range strings.Fields(s) {
    a[word] += 1
  }
  
  return a
}

func main() {
  wc.Test(WordCount)
}
