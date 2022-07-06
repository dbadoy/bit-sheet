package main

func main() {
  const uintSize = 32 << (^uint(0) >> 63)
  // if uint16 | uint32 
  // ^uint(0) >> 63 : 0
  // uintSize = 32
  
  // if uint64
  // ^uint(0) >> 63 : 1
  // uintSize = 32 << 1
  // uintSize = 64
}
