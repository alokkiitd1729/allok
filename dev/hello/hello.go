package main

import (
	"fmt"
)

type test struct {
	ok  int
	ok1 int
}

func NewTest(v int) *test {
	return &test{
		ok:  v,
		ok1: v * v,
	}
}
// Adding a comment to see this in a new commit.
func greetings(name string) {
	fmt.Printf("Hi there!! %s", name)
}

func main() {
	tmp := NewTest(4)
	fmt.Println(tmp)
	fmt.Printf("%p\n", tmp)
	fmt.Println(tmp.ok)
	fmt.Println(*tmp)
	fmt.Println(&tmp)
}
