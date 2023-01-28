package main

import (
	"fmt"
	"math"
)

func main() {
	b := []bool{}

  for i := 0; i < 101; i++{
    if(i < 99){
      b = append(b, false)
    } else {
      b = append(b, true)
    }
  }

	fmt.Println(CrystalBalls(b))
}

func CrystalBalls(b []bool) int {
  //Jumpt the sqrrt of the length
	j := int(math.Sqrt(float64(len(b))))
  //Start I at the position j if it's true we jump back and iterate through
	i := j

	for i < len(b) {
		fmt.Println(b[i], i)
		if b[i] {
			break
		}
		i += j
	}
  //move back and then iterate through to find the first true position
	i -= int(j)

	for k := 0; k < j && i < len(b); {
		k += 1
		i += 1
		fmt.Println(b[i], i, k)
		if b[i] {
			return i
		}

	}

	return -1

}
