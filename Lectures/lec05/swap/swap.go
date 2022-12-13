// swap/swap.go
package main

func swap(pa, pb *int) {
	*pa, *pb = *pb, *pa
}
