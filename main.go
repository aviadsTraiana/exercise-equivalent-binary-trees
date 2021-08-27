package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var recWalk func(t *tree.Tree)
	recWalk = func(t *tree.Tree) {
		if t != nil {
			recWalk(t.Left)
			ch <- t.Value
			recWalk(t.Right)
		}
	}
	recWalk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		c1, ok1 := <-ch1
		c2, ok2 := <-ch2
		if ok1 && ok2 {
			if c1 != c2 {
				return false
			}
		} else {
			return ok1 == ok2
		}
	}

}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for c := range ch {
		fmt.Print(c)
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
