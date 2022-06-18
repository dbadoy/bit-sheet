package main

import "fmt"

const (
	A_FLAG = 0x01
	B_FLAG = 0x02
	C_FLAG = 0x04
	D_FLAG = 0x08
	E_FLAG = 0x010
)

// present by bit shift
const (
	A_FLAG_SHIFT = (1 << 0)
	B_FLAG_SHIFT = (1 << 1)
	C_FLAG_SHIFT = (1 << 2)
	D_FLAG_SHIFT = (1 << 3)
	E_FLAG_SHIFT = (1 << 4)
)

func main() {
	fmt.Println(A_FLAG)
	fmt.Println(B_FLAG)
	fmt.Println(C_FLAG)
	fmt.Println(D_FLAG)
	fmt.Println(E_FLAG)

	fmt.Println(A_FLAG_SHIFT)
	fmt.Println(B_FLAG_SHIFT)
	fmt.Println(C_FLAG_SHIFT)
	fmt.Println(D_FLAG_SHIFT)
	fmt.Println(E_FLAG_SHIFT)
	
	A_FLAG_CHECKER := func(v int) bool {
		return (v & A_FLAG) > 0
	}

	// true
	fmt.Println(A_FLAG_CHECKER(A_FLAG))
	// true
	fmt.Println(A_FLAG_CHECKER(A_FLAG_SHIFT))
	// false
	fmt.Println(A_FLAG_CHECKER(E_FLAG))
}
