package main

import (
	"crypto/aes"
	"fmt"
)

// Bob the evaluator

func evaluateGarbledCircuit(inputs [][]byte, gates []Gate) []byte {
	n := len(inputs) // number of inputs
	m := len(gates)  // number of gates

	// array of signals have a size of n+m
	signals := make([][]byte, m+n)

	// setup inputs signals
	for i := 0; i < n; i++ {
		signals[i] = inputs[i]
		fmt.Printf("input %d=%x\n", i, signals[i][:2])
	}

	// add code below to evaluate the gates
	for _, gate := range gates {
		a := signals[gate.in0]
		b := signals[gate.in1]
		sa := (a[0] & 0x80) >> 7
		sb := (b[0] & 0x80) >> 7
		sasb := sa*2 + sb
		g, _ := aes.NewCipher(append(a, b...))
		ug := make([]byte, 16)
		g.Decrypt(ug, gate.table[sasb])
		signals[gate.out] = ug
	}

	// the last signal is the output
	return signals[n+m-1]
}
