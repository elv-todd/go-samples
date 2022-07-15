package main

import "golang.org/x/tour/tree"
import "fmt"

//
// Solution to https://go.dev/tour/concurrency/8
// see https://pkg.go.dev/golang.org/x/tour/tree#Tree

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var tres1, tres2 [10]int
	for i := 0; i < 10; i++ {
		tres1[i] = <-ch1
		tres2[i] = <-ch2
	}
	fmt.Printf("tree1 %v -> %v\n", t1, tres1)
	fmt.Printf("tree2 %v -> %v\n", t2, tres2)

	return tres1 == tres2
}

func main() {
	ch := make(chan int, 100)
	go Walk(tree.New(14), ch)
	for i := 0; i < 10; i++ {
		fmt.Printf("walked %d ", <-ch)
	}
	fmt.Printf("done walk\n")

	t1 := tree.New(13)
	t2 := tree.New(13)
	fmt.Printf("same? (should be yes) %v\n", Same(t1, t2))
	t2 = tree.New(12)
	fmt.Printf("same? (should be no) %v\n", Same(t1, t2))
}
