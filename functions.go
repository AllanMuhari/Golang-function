package main

import (
	"fmt"
)

func add(x, y int) (int, int) {

	return x + y, x - y
}

func main() {

	ans1, ans2 := add(10, 6)
	fmt.Println(ans1, ans2)
	// a function inside another function
	test := func(x int) int {
		return x * 2
	}(5)
	fmt.Println(test)
}
