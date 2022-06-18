package main

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
  
}
