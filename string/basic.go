package main

import "fmt"

func main() {
	str := "Hello, world!"
	r := []rune(str)

	Convert(r)

	indices := IndexOf(r, 'l')
	// [2 3 10]
	fmt.Println(indices)
	for _, v := range indices {
		// l
		// l
		// l
		fmt.Println(string(str[v]))
	}

}

func Convert(r []rune) {
	lower := []rune{}
	upper := []rune{}
	invert := []rune{}

	for _, v := range r {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			lower = append(lower, (v | ' '))
			upper = append(upper, (v & '_'))
			invert = append(invert, (v ^ ' '))
		} else {
			lower = append(lower, v)
			upper = append(upper, v)
			invert = append(invert, v)
		}
	}
	// hello, world!
	fmt.Println(string(lower))
	// HELLO, WORLD!
	fmt.Println(string(upper))
	// hELLO, WORLD!
	fmt.Println(string(invert))
}

func IndexOf(r []rune, char rune) []int {
	// find index
	indices := []int{}
	for i, v := range r {
		if (v ^ char) == 0 {
			indices = append(indices, i)
		}
	}
	return indices
}
