package main

import "fmt"

func assign() {
	a := []int{0, 1, 2, 3, 4}
	b := a

	b[0] = 100

	fmt.Printf("after assign: a=%v, b=%v\n", a, b)
}
