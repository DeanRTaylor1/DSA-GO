package main

import (
	"fmt"
)

func main() {
  s := []int64{}

  for i := 0; i < 100001; i++{
    s = append(s, int64(i))
 }

  fmt.Printf("%v", BinarySearch(s, 3))
  }

func BinarySearch(s []int64, i int64) bool {
  l := 0
  h := len(s)
	fmt.Println(l)
  for l < h {
    m := (l + (h - l) / 2) 
    v := s[m]
    fmt.Println(l, h, m)
    if(v == i){
      return  true
    } else if v > i {
      h = m
    } else {
      l = m + 1
    }
  }

  return false

}
