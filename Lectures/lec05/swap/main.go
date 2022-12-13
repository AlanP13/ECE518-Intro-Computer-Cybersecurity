// swap/main.go
package main

import "fmt"

func main() {
	var a int = 1
	b := 2
	fmt.Printf("before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap: a = %d, b = %d\n", a, b)
}
