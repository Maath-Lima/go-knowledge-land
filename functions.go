// If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.

package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func rectangleArea(l1, l2 int) (area int, perimeter int) {
	area = l1 * l2
	perimeter = 2 * (l1 + l2)

	return
}

func update(age *int, name *string) {
	*age = *age + 5
	*name = *name + " Doe"

	return
}

type First func(int) int

func partialSum(x int) First {
	return func(y int) int {
		return sum(x, y)
	}
}

func main() {
	fmt.Println(sum(32, 27))

	var a, p int = rectangleArea(2, 4)
	fmt.Println(a)
	fmt.Println(p)

	age := 25
	name := "Jhon"

	// Sample using closures functions

	buildText := func(description, name string, age int) {
		fmt.Println(description, name, age)
	}

	buildText("Before update:", name, age)

	update(&age, &name)

	buildText("After update:", name, age)

	// Sample anonymous function

	fmt.Printf("100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	// HOF
	partial := partialSum(3)
	fmt.Println(partial(7))
}
