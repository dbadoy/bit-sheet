package main

func Average() {
  x := int32(8669)
  y := int32(32318)
  
  // (x + y) / 2
  // 20493.5
  
  // 20493
  // TODO: add comment
  fmt.Println((x + y) >> 1)
  
  // floor
  // 20493
  fmt.Println((x & y) + ((x ^ y) >> 1))
  
  // ceiling
  // 20494
  fmt.Println((x | y) - ((x ^ y) >> 1))
}
