package main

import "fmt"

func mycopy() {
	a := []int{0, 1, 2, 3, 4}

	b := make([]int, len(a))
	copy(b, a)

	b[0] = 100

	fmt.Printf("after copy: a=%v, b=%v\n", a, b)
}
