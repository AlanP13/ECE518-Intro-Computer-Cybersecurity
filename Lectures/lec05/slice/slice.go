// slice/slice.go
package main

import "fmt"

func main() {
	var a [10]int
	s := make([]int, 0)
	for i := 0; i < 10; i++ {
		a[i] = i
		s = append(s, i*i)
	}
	for i, val := range s {
		fmt.Printf("s[%d]=%d=%d*%d\n", i, val, a[i], a[i])
	}
	assign()
	mycopy()
	myappend()
	slicing()
}
