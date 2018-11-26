package main

// Solution for: https://tour.golang.org/methods/23

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
  n, err = rot.r.Read(b)
  for i,v := range b {
    if (v >= 'a' && v < 'n') || (v >= 'A' && v < 'N') {
      b[i] = v + 13
    } else if (v > 'm' && v <= 'z') || (v > 'M' && v <= 'Z') {
      b[i] = v - 13
    }
  }
  return
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
