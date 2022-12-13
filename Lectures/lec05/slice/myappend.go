package main

import "fmt"

func myappend() {
	a := []int{100}

	// don't do this
	for i := 0; i < 10; i++ {
		b := a
		a = append(a, i)
		b[0]++
		fmt.Printf("append %d: a=%v, b=%v\n", i, a, b)
	}
}
