package main

import (
	"fmt"
)

// if and switch accept an initialization statement
func main() {
	if test := false; test {
		fmt.Println("Hello If")
	}
}
