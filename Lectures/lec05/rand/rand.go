// rand/rand.go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	d := rand.Float64()
	if d < 0.4 {
		fmt.Println("Win!")
	} else if d > 0.6 {
		fmt.Println("Lose!")
	} else {
		fmt.Println("Tie!")
	}
}
