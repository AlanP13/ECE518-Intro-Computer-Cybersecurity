package main

import "fmt"

func slicing() {
	a := []int{0, 1, 2, 3, 4}
	b := a[1:3]
	c := a[:len(a)-1]
	d := a[2:]

	fmt.Printf("a=%v, b=%v, c=%v, d=%v\n", a, b, c, d)
}
