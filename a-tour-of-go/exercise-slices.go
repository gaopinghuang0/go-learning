package main

// solution for: https://tour.golang.org/moretypes/18

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  grid := make([][]uint8, dy);
  for i := range grid {
    row := make([]uint8, dx)
    for j := range row {
      row[j] = uint8((i+j)/2)
    }
    grid[i] = row
  }
  return grid
}

func main() {
  pic.Show(Pic)
}
