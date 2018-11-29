package main

// solution for: https://tour.golang.org/concurrency/8

import (
  "fmt"
  "golang.org/x/tour/tree"
)

// Credit: https://stackoverflow.com/a/12224635/4246348
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  // it is an idiomatic way to defer close ch here, rather than close at the end
  defer close(ch)  // <- closes the channel when this function returns
  var walk func(t *tree.Tree)
  walk = func(t *tree.Tree) {
    if t == nil { return }
    // must be in-order traversal
    walk(t.Left)
    ch <- t.Value
    walk(t.Right)   
  }
  walk(t)
}

// Credit: https://stackoverflow.com/a/40534404/4246348
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
    for {
        v1,ok1 := <- ch1
        v2,ok2 := <- ch2
        if v1 != v2 || ok1 != ok2 {
            return false
        }
        if !ok1 {
            break
        }
    }

    return true
}

func testWalk() {
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for i := range ch {
    fmt.Println(i)
  }
}

func testSame() {
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(1), tree.New(2)))
}

func main() {
  testWalk()
  testSame()
}
