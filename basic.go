package main

import "fmt"

func main() {
	// a << n (a * 2^n)
	// a >> n (a / 2^n)

	x := int32(1)

	// 000000000100000
	fmt.Printf("%b\n", (x << 5))

	// 000000000011111
	fmt.Printf("%b\n", ((x << 5) - 1))

	// check 2^k
	// if zero, it's true
	fmt.Println((x << 5) & ((x << 5) - 1))
	
	// check odd number
	fmt.Println((79 & 1) == 1)

	// check sign
	y := -91
	z := 30
	// false
	fmt.Println((y ^ z) >= 0)

	// abs
	// -71
	fmt.Println(71 ^ -1 + 1)
	// 101
	fmt.Println(-101 ^ -1 + 1)
}
