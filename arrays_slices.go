package main

import (
	"fmt"
)

func arrays() {
	// grades := [3]int{97, 98, 99}
	grades := [...]int{97, 98, 99}

	fmt.Println(grades)
	fmt.Println(len(grades))

	// Matrix
	identityMatrix := [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	fmt.Println(identityMatrix)

	// Pointers
	a := [...]int{1, 2, 3}
	b := &a

	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}

func slices() {
}

func main() {
	arrays()
}
