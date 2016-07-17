package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        Walk(t.Right, ch)
    }
    return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)

    defer func() {
        for range ch1 {
        }
        for range ch2 {
        }
    }()

    go func() {
        Walk(t1, ch1)
        close(ch1)
    }()

    go func() {
        Walk(t2, ch2)
        close(ch2)
    }()

    for {
        i, ok1 := <-ch1
        j, ok2 := <-ch2

        if i != j || ok1 == false && ok2 == true || ok1 == true && ok2 == false {
            return false
        } else {
            if ok1 == false {
                return true
            }
        }
    }
}

func main() {
    l := Same(tree.New(1), tree.New(1))
    m := Same(tree.New(1), tree.New(2))
    fmt.Println(l)
    fmt.Println(m)
}
